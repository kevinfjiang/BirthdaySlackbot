package main
// TODO write a shema for the database (ON PAPER), set up cloud watch, and that's all
import (
        "fmt"
		"log"
		"context"
		"errors"

		"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot"
		"github.com/kevinfjiang/BirthdayServer/src/DB"

		"github.com/aws/aws-lambda-go/lambda"
)

 var Cred *BirthdayBot.Creds

func init(){
	Cred = BirthdayBot.GetCreds() // New connection esstablished everytime.
	// TODO write more tests to ensure functions are all valid, consider using concurrency if need be
}

func NotFound(Type string) error {
	return errors.New(fmt.Sprintf("MSGEvent Type not found %s", Type))
}

func HandleRequest(ctx context.Context, event DB.MSGEvent) (string, error) {
	log.Printf("[INFO] Received event %s", event.Type)

	var reply string 
	var err error = nil

	switch event.Type{
	case "Wish":
		reply = Cred.DB_Connect.MessageHandle(&event)
	case "DailyPing":
		reply = Cred.SendDailyMSG(event.SendPM) // All it does is check the typpe and moves from there
	case "GenWeb":
		reply = ""
	default:
		reply = ""
		err = NotFound(event.Type)
	} 
	return reply, err
}



func main() {
	lambda.Start(HandleRequest)
}