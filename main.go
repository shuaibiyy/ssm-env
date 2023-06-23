package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/joho/godotenv"
	"os"
)

func getSsmParam(ssmSvc *ssm.SSM, keyName string) string {
	param, err := ssmSvc.GetParameter(&ssm.GetParameterInput{
		Name:           &keyName,
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		panic(errors.New(
			fmt.Sprintf("Error: %v\nParameter: %s\n", err.Error(), keyName)))
	}

	value := *param.Parameter.Value
	return value
}

func substituteSsmParams(ssmSvc *ssm.SSM, env map[string]string) map[string]string {
	var substEnv = make(map[string]string)

	for k, v := range env {
		if strings.HasPrefix(v, "ssm:") {
			newVal := getSsmParam(ssmSvc, v[4:])
			substEnv[k] = newVal
		} else {
			substEnv[k] = v
		}
	}

	return substEnv
}

func main() {
	var region string
	var srcPath string

	flag.StringVar(&region, "region", os.Getenv("AWS_REGION"), "AWS region")
	flag.StringVar(&srcPath, "src", os.Getenv("SRC_PATH"), "Source env file path")
	flag.Parse()

	if region == "" || srcPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	sess, err := session.NewSession(
		&aws.Config{Region: aws.String(region)},
	)
	if err != nil {
		panic(err)
	}

	ssmSvc := ssm.New(sess)

	var env = make(map[string]string)
	env, err = godotenv.Read(srcPath)
	if err != nil {
		panic(err)
	}

	substEnv := substituteSsmParams(ssmSvc, env)
	content, _ := godotenv.Marshal(substEnv)
	fmt.Printf("%v\n", content)
}
