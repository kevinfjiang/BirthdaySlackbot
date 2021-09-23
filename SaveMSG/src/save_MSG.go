package main
// TODO write a shema for the database (ON PAPER), set up cloud watch, and that's all
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