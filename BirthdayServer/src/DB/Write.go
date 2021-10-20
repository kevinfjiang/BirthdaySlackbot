package DB

import (
	"fmt"
	"log"
	// "time"

	"github.com/kevinfjiang/BirthdayServer/src/MSGEvent"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)


func process_request(message *MSGEvent.Request) *PMessage{
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

func (DB DBConnect) MessageHandle(message *MSGEvent.Request) error{
    request := process_request(message)
	if err := DB.write(request); err != nil{
		log.Fatalf("[ERROR] %s", err)
	}
	return nil

}
