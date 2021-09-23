package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinfjiang/slackBirthdayBot/src/googleSheets"
	"github.com/kevinfjiang/slackBirthdayBot/src/slackMSG"
	// "github.com/slack-go/slack"
)

func main(){
	// fmt.Println(os.Getenv("GOOGLE_API_JSON"))
	if err := godotenv.Load(os.Getenv("ENV_PATH")); err==nil {
		api := slackMSG.New_SlackAPI(os.Getenv("SLACKBOT_TOKEN"))
		FB := googleSheets.GetTable(os.Getenv("GOOGLE_API_JSON"), os.Getenv("GOOGLE_SHEETS_ID"), "B:E", api)
		PreBdays, Bdays := googleSheets.Find_BDAYS(FB)
		googleSheets.Prep_BDAY_MSG(PreBdays, Bdays, FB, api)
	}
}