package DB

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DBConnect struct {
	TableName string
	*dynamodb.DynamoDB 
}

func Get_DB_Connect() DBConnect {
	log.Print("[INFO] Creating connection to DB")
	sess := session.Must(session.NewSession()) //TODO establish legit connection to DB, rn not great! Probs as an env variable
	
	sess.Handlers.Send.PushFront(func(r *request.Request) {
		log.Printf("[INFO] Request: %s/%v, Params: %s",
			r.ClientInfo.ServiceName, r.Operation, r.Params)
	})

	return DBConnect{"Table", dynamodb.New(sess)} 
}
