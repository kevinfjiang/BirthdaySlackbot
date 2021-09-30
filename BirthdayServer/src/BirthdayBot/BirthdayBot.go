package BirthdayBot

import (
	"os"

	"github.com/kevinfjiang/BirthdayServer/src/DB"
	"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot/Sheets"
	"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot/SlackMSG"
)

type Creds struct{
	SLACKBOT_TOKEN string
	GOOOGLE_API_JSON string
	GOOGLE_SHEETS_ID string
	DB_TOKEN string
}

func GetCreds()*Creds{
	return &Creds{
		SLACKBOT_TOKEN:   os.Getenv("SLACKBOT_TOKEN"),
		GOOOGLE_API_JSON: os.Getenv("GOOGLE_API_JSON"),
		GOOGLE_SHEETS_ID: os.Getenv("GOOGLE_SHEETS_ID"),
	}

}

func GetBDAYPeople(Cred *Creds) []*Sheets.Staff{
	api := SlackMSG.New_SlackAPI(os.Getenv("SLACKBOT_TOKEN"))
	FB := Sheets.GetTable(os.Getenv("GOOGLE_API_JSON"), os.Getenv("GOOGLE_SHEETS_ID"), "B:E", api)
	_, Bdays := Sheets.Find_BDAYS(FB)

	var StaffBdays []*Sheets.Staff // Type correction
	for _, staffmember := range(Bdays){
		StaffBdays = append(StaffBdays, staffmember.(*Sheets.Staff))
	}


	return StaffBdays
}

func SendDailyMSG(Cred *Creds) {
	api := SlackMSG.New_SlackAPI(os.Getenv("SLACKBOT_TOKEN"))
	FB := Sheets.GetTable(os.Getenv("GOOGLE_API_JSON"), os.Getenv("GOOGLE_SHEETS_ID"), "B:E", api)
	PreBdays, Bdays := Sheets.Find_BDAYS(FB)

	dbConnect := DB.Get_DB_Connect()

	Sheets.Send_BDAY_Private_MSG(Bdays, dbConnect, api)
	Sheets.Prep_BDAY_MSG(PreBdays, Bdays, FB, api)
}
