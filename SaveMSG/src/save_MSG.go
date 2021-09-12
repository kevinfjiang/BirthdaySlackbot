package main

import (
        // "fmt"
        // "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func handleRequest () (string, error) {
    return "Hello from Go!", nil
}

func main() {
	lambda.Start(handleRequest)
}