package DB

import (
	"fmt"
	"log"
	// "time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type MSGEvent struct {
	Type	 		string `json:"Type of MSG"`
	BirthdayPerson	string `json:"The Bday Pal ID"`
	SenderPerson 	string `json:"The Sender's ID"`

	Message 		string `json:"Message for user"`

	SendPM			bool   `json:"Whether to send private messsages"`

}

func process_request(message *MSGEvent) *PMessage{
	// TODO Lot of processing necessary
	return nil
}

func (DB DBConnect) write(message *PMessage) error { // Incorporate a Time to live
	row, err := dynamodbattribute.MarshalMap(*message)
	if err != nil {
		log.Fatalln(fmt.Sprintf("[FATAL] failed to DynamoDB marshal Record, %v", err))
	}

	_, err = DB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(DB.TableName),
		Item:      row,
	})
	if err != nil {
		log.Fatalln(fmt.Sprintf("[FATAL] Failed to put Record to DynamoDB, %v", err))
	}
	return nil
}

func (DB DBConnect) MessageHandle(message *MSGEvent) string{
    request := process_request(message)
	if err := DB.write(request); err != nil{
		log.Fatalf("[ERROR] %s", err)
	}
	return ""

}
