package main

import (
	"context"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

func main() {
	parser := argparse.NewParser("signed-s3-url-generator", "Generates signed urls for S3 uploads.")
	bucketName := parser.String("b", "bucket", &argparse.Options{Required: true, Help: "Name of the bucket to create the signed url for."})
	objectKey := parser.String("k", "key", &argparse.Options{Required: true, Help: "Key in the bucket to create the signed url for."})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client := s3.NewFromConfig(cfg)
	checkBucket(client, bucketName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	presignClient := s3.NewPresignClient(client)
	request, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(*bucketName),
		Key:    aws.String(*objectKey),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(request.URL)
}

func checkBucket(client *s3.Client, bucketName *string) {
	_, err := client.HeadBucket(context.TODO(), &s3.HeadBucketInput{Bucket: bucketName})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
