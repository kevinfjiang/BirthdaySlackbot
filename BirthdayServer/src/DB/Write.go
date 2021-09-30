package DB

import (
	// "fmt"
	"log"
	// "time"
	"net/http"

	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func process_request(r *http.Request) *PMessage{
	return nil
}

func (DB DBConnect) write(*PMessage) error {
	return nil
}

func (DB DBConnect) PMHandler(w http.ResponseWriter, r *http.Request) {
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
