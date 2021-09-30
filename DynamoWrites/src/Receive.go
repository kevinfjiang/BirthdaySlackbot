package main
// TODO write a shema for the database (ON PAPER), set up cloud watch, and that's all
import (
        "fmt"
		"log"
        "net/http"

		"github.com/kevinfjiang/slackBirthdayBot/DynamoWrites/src/DBWrite"
)


func main() {
	dyno := DBWrite.Get_DB_Connect()

	srv := &http.Server{
        Addr:         ":8080",
        Handler:      http.HandlerFunc(dyno.pmHandler),
    }
	defer srv.Close()
	if err_http := srv.ListenAndServe(); err_http != nil {
		log.Fatal("[ERROR] Server Crashed: ", err_http)
	}
}