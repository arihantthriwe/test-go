package main

import (
	"test-go/lambda"

	aws_lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	aws_lambda.Start(lambda.Handler)
}
