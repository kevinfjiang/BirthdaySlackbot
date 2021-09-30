package DB

import (
	// "fmt"
	"log"
	// "time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DBConnect struct {
	*dynamodb.DynamoDB 
}

func Get_DB_Connect() DBConnect {
	log.Print("[INFO] Creating connection to DB")
	sess := session.Must(session.NewSession())
	
	sess.Handlers.Send.PushFront(func(r *request.Request) {
		log.Printf("[INFO] Request: %s/%s, Params: %s",
			r.ClientInfo.ServiceName, r.Operation, r.Params)
	})

	return DBConnect{dynamodb.New(sess)} // Double check this is valids
}