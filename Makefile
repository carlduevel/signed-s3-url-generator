SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --always-make
# comment out next line for debugging
MAKEFLAGS += --silent

ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

fmt: ## Fmt all files
> go fmt ./...

clean: ## Remove temporary files
> rm -rf bin

bin/signed-s3-url-generator_amd64: *.go ## create binary for Linux
> GOOS=linux GOARCH=amd64 go build -o $@

bin/signed-s3-url-generator_darwin: *.go ## create binary for Mac on Intel
> GOOS=darwin GOARCH=amd64 go build -o $@

bin/signed-s3-url-generator_darwin_arm: *.go ## create binary for Mac on ARM
> GOOS=darwin GOARCH=arm64 go build -o $@

bin/signed-s3-url-generator.exe: *.go ## create binary for Windows
> GOOS=linux GOARCH=amd64 go build -o $@

bin/all: bin/signed-s3-url-generator_amd64 bin/signed-s3-url-generator_darwin bin/signed-s3-url-generator_darwin_arm bin/signed-s3-url-generator.exe ##create binaries for all platforms

help:
> egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help