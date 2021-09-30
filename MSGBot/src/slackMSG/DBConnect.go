package SlackMSG

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DBConnect struct {
	*dynamodb.DynamoDB 
}

type Message struct { // TODO define Message
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

func (DB DBConnect) Get_MSG(birthday map[string]interface{}) []*Message {
	log.Printf("[INFO] Reading DB for %s", birthday["Name"].(string))
	result, err := DB.Query(&dynamodb.QueryInput{
		TableName: aws.String("Bday_Messages"),
		KeyConditionExpression: aws.String(fmt.Sprint(
			"partitionKeyName = :%s_%s",
			birthday["Slack_ID"],
			time.Now().Year())),
	})
	log.Printf("[INFO] Reading DB finished for %s", birthday["Name"].(string))

	if err != nil {
		log.Fatalln("[ERROR] Got error calling GetItem: %s", err)
	}

	if *(result.Count) == int64(0) {
		log.Printf("[INFO] Found Messages Found for %s", time.Now().Format("January 2, 2006"))
		return nil
	}

	var Messages []*Message
	for i, _ := range result.Items {
		response := Message{} // Consider multiple rows, just iterate through the rows
		err = dynamodbattribute.UnmarshalMap(result.Items[i], &response)
		if err != nil {
			log.Printf("[WARNING] Failed to unmarshal Record: %v", err)
		}
		Messages = append(Messages, &response)

	}

	return Messages

}
