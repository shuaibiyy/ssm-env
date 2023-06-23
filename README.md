# ssm-env 

Substitute SSM parameters in env files.

If you have an env file [`example.env.tmpl`](./example.env.tmpl) with the following content:
```
SECRET=ssm:/com/team/svc
STATIC=bar
```
Running `ssm-env -src example.env.tmpl` will substitute all entries that begin with the prefix `ssm:`, resulting in the following output:
```
SECRET=foo
STATIC=bar
```
assuming the value of `/com/team/svc` in AWS Parameter Store is `foo`.

## Installation

Download a binary from the [releases page](https://github.com/shuaibiyy/ssm-env/releases).

## Usage
```
ssm-env
  -region string
        AWS region
        Default: AWS_REGION environment variable 
  -src string
        Source env file path
        Default: SRC_PATH environment variable
```

## Development
```
brew install dep
dep ensure
go run main.go
```

## Build
```
GOOS=darwin go build
GOOS=linux go build
```

## Release
This project uses [GoReleaser](https://goreleaser.com/).
```
export GITHUB_TOKEN=`YOUR_GH_TOKEN`
goreleaser
```
