# signed-s3-url-generator

Tiny tool to generate [presigned AWS S3
URLs](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ShareObjectPreSignedURL.html).


## Example usage

You need to upload data from a machine that is not authorized to access an S3
bucket (and should not do so otherwise). So on a machine with authorization create a url using the
presigned url:
```
signed-s3-url-generator -b mybucket -k 'key/to/save/this/file/under/example.txt'
```
On the machine without S3 authorization run:
```
curl <url generated earlier> --upload-file example.txt
```


