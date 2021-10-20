package main
// TODO write a shema for the database (ON PAPER), set up cloud watch, and that's all
import (
		"log"
		"context"

		"github.com/kevinfjiang/BirthdayServer/src/MSGEvent"
		"github.com/kevinfjiang/BirthdayServer/src/DB"
		

		"github.com/aws/aws-lambda-go/lambda"
)

 var DB_Connect DB.DBConnect

func init(){
	DB_Connect = DB.Get_DB_Connect() // New connection esstablished everytime.
	// TODO write more tests to ensure functions are all valid, consider using concurrency if need be
	// TODO set up logger
}


func HandleRequest(ctx context.Context, event MSGEvent.Request) (MSGEvent.Response, error) {
	var reply MSGEvent.Response 
	var err error = nil

	log.Printf("[INFO] Received event %s", event.Type)
	switch event.Type{
	case "Wish":
		err = DB_Connect.MessageHandle(&event)
		reply = MSGEvent.GenSuccessfulMessageSentResponse()
	default:
		reply = MSGEvent.GenNotFoundResponse()
	} 

	if err!=nil {
		reply = MSGEvent.GenErrorResponse()
	}
	return reply, err // TODO try not to return nil!
}



func main() {
	lambda.Start(HandleRequest)
}