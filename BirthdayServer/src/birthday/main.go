package main
// TODO write a shema for the database (ON PAPER), set up cloud watch, and that's all
import (
		"log"
		"context"

		"github.com/kevinfjiang/BirthdayServer/src/MSGEvent"
		"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot"

		"github.com/aws/aws-lambda-go/lambda"
)

 var Cred *BirthdayBot.Creds

func init(){
	Cred = BirthdayBot.GetCreds() // New connection esstablished everytime.
	// TODO write more tests to ensure functions are all valid, consider using concurrency if need be
	// TODO set up logger
}

func HandleCall(ctx context.Context, event MSGEvent.Request) (MSGEvent.Response, error) {
	var reply MSGEvent.Response 
	var err error = nil

	log.Printf("[INFO] Received event %s", event.Type)
	reply = MSGEvent.GenNotFoundResponse()
	switch event.Type{
	case "DailyPing":
		err = Cred.SendDailyMSG(event.SendPM) // All it does is check the type and moves from there
	default:
		reply = MSGEvent.GenNotFoundResponse()
	} 

	if err!=nil {
		reply = MSGEvent.GenErrorResponse()
	}
	return reply, err // TODO try not to return nil!
}



func main() {
	lambda.Start(HandleCall)
}