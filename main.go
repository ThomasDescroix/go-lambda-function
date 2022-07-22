package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type LambdaResponse struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(event LambdaEvent) (LambdaResponse, error) {
	return LambdaResponse{Message: fmt.Sprintf("Hello %s, you are %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
