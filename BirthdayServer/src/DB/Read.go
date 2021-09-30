package DB

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type PMessage struct { // TODO define Message
}

func (DB DBConnect) Get_MSG(birthday map[string]interface{}) []*PMessage {
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

	var Messages []*PMessage
	for i, _ := range result.Items {
		response := PMessage{} // Consider multiple rows, just iterate through the rows
		err = dynamodbattribute.UnmarshalMap(result.Items[i], &response)
		if err != nil {
			log.Printf("[WARNING] Failed to unmarshal Record: %v", err)
		}
		Messages = append(Messages, &response)

	}

	return Messages

}
