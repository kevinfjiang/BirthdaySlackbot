package main

import(
	// "fmt"

	"os"
	"github.com/kevinfjiang/slackBirthdayBot/src/googleSheets"
	"github.com/kevinfjiang/slackBirthdayBot/src/slackMSG"
	// "github.com/slack-go/slack" 
)

func main(){
	// fmt.Println(os.Getenv("GOOGLE_API_JSON"))
	api := slackMSG.New_SlackAPI(os.Getenv("SLACKBOT_TOKEN"))
	FB := googleSheets.GetTable(os.Getenv("GOOGLE_API_JSON"), os.Getenv("SLACKBOT_TOKEN"), "B:E", api)
	PreBdays, Bdays := googleSheets.Find_BDAYS(FB)
	googleSheets.Prep_BDAY_MSG(PreBdays, Bdays, FB, api)

}