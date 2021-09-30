package DBWrite


import (
	"fmt"
	"log"
	"time"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DBConnect struct {
	*dynamodb.DynamoDB 
}

type PMessage struct { // TODO define Message
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

func (DB DBConnect) write(PMessage)error{

}

func (DB DBConnect) pmHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "Post" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

	// if err := r.ParseForm(); err != nil {
    //     fmt.Fprintf(w, "ParseForm() err: %v", err)
    //     return
    // }

    request := process_request(r)
	if err := DB.write(request); err != nil{
		log.Fatal("[ERROR] %v", err)
	}

}