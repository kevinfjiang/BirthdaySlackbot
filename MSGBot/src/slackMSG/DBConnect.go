package SSlackMSG

import (
	"log"
	"request"

	"github.com/aws"
	"github.com/aws/aws-sdk-go"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	
)

type DBConnect struct{
	*dynamodb.DynamoDB // Do i name it or nah??
}

type Message struct {

}


func Get_DB_Connect() *DBConnect{
	sess := session.Must(session.NewSession())
	sess.Handlers.Send.PushFront(func(r *request.Request) {
		// Log every request made and its payload
		log.Printf("[INFO] Request: %s/%s, Params: %s",
			r.ClientInfo.ServiceName, r.Operation, r.Params)
	})
	return dynamodb.New(sess) // Double check this is valids
}

func (DB DBConnect) Get_MSG(birthday map[string]interface{}) []*Message{
	result, err := DB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(movieYear),
			},
			"Title": {
				S: aws.String(movieName),
			},
		},
	})
	
	if err != nil {
		logger.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		msg := "Could not find '" + *title + "'"
		return nil
	}
	
	var Messages []*Message
	for i, _ := range(result.Item){
		response := Message{} // Consider multiple rows, just iterate through the rows
		err = dynamodbattribute.UnmarshalMap(result.Item[i], &response)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		}
		Messages = append(Messages, &response)

	}

	return Messages
	
}