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



func NotFound(Type string) error {
	return errors.New(fmt.Sprintf("MSGEvent Type not found %s", Type))
}

func HandleRequest(ctx context.Context, event DB.MSGEvent) (string, error) {
	log.Printf("[INFO] Received event %s", event.Type)
	Cred := BirthdayBot.GetCreds() // New connection esstablished everytime.
	var reply string = ""
	var err error = nil

	switch event.Type{
	case "Wish":
		reply = Cred.DB_Connect.MessageHandle(&event)
	case "DailyPing":
		reply = Cred.SendDailyMSG() // All it does is check the typpe and moves from there
	case "GenWeb":
		reply = ""
	default:
		err = NotFound(event.Type)
	} 
	return reply, err
}



func main() { // Rewrite thiss, some how include multiple handle funcs
	// Figure out integration of the outter iteraction with the DB and BirthdayBot Interaction
	
	lambda.Start(HandleRequest)
}