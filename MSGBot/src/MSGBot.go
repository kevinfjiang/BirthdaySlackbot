package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinfjiang/slackBirthdayBot/src/Sheets"
	"github.com/kevinfjiang/slackBirthdayBot/src/SlackMSG"
)

func main(){
	// fmt.Println(os.Getenv("GOOGLE_API_JSON"))
	if err := godotenv.Load(os.Getenv("ENV_PATH")); err==nil {
		api := SlackMSG.New_SlackAPI(os.Getenv("SLACKBOT_TOKEN"))
		FB := Sheets.GetTable(os.Getenv("GOOGLE_API_JSON"), os.Getenv("GOOGLE_SHEETS_ID"), "B:E", api)
		PreBdays, Bdays := Sheets.Find_BDAYS(FB)

		dbconnect := SlackMSG.Get_DB_Connect()

		Sheets.Send_BDAY_Private_MSG(Bdays, dbConnect, api)
		Sheets.Prep_BDAY_MSG(PreBdays, Bdays, FB, api)
	}
}