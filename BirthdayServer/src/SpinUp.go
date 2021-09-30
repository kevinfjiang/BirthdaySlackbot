package main
// TODO write a shema for the database (ON PAPER), set up cloud watch, and that's all
import (
        // "fmt"
		"log"
        "net/http"

		"github.com/kevinfjiang/BirthdayServer/src/DB"
		"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot"
)


func main() {
	_ = BirthdayBot.GetCreds() // Figure out integration of the outter iteraction with the DB and BirthdayBot Interaction
	dyno := DB.Get_DB_Connect()

	srv := &http.Server{
        Addr:         ":8080",
        Handler:      http.HandlerFunc(dyno.PMHandler),
    }
	
	defer srv.Close()
	if err_http := srv.ListenAndServe(); err_http != nil {
		log.Fatal("[ERROR] Server Crashed: ", err_http)
	}
}