package BirthdayBot

import (
	"os"

	"github.com/kevinfjiang/BirthdayServer/src/DB"
	"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot/Sheets"
	"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot/SlackMSG"
)

type Creds struct{
	SLACKBOT_TOKEN string
	GOOGLE_API_JSON string
	GOOGLE_SHEETS_ID string

	DB_Connect DB.DBConnect // Write DB errors one of these days!!
}

func GetCreds()*Creds{
	return &Creds{
		SLACKBOT_TOKEN:		os.Getenv("SLACKBOT_TOKEN"),
		GOOGLE_API_JSON:  	os.Getenv("GOOGLE_API_JSON"),
		GOOGLE_SHEETS_ID: 	os.Getenv("GOOGLE_SHEETS_ID"),

		DB_Connect:		  	DB.Get_DB_Connect(),
	}

}

func (Cred *Creds) GetBDAYPeople() []*Sheets.Staff{
	api := SlackMSG.New_SlackAPI(Cred.SLACKBOT_TOKEN)
	FB := Sheets.GetTable(Cred.GOOGLE_API_JSON, Cred.GOOGLE_SHEETS_ID, "B:E", api)
	
	_, Bdays := Sheets.Find_BDAYS(FB)

	var StaffBdays []*Sheets.Staff // Type correction
	for _, staffmember := range(Bdays){
		StaffBdays = append(StaffBdays, staffmember.(*Sheets.Staff))
	}

	return StaffBdays
}

func (Cred *Creds) SendDailyMSG() string{
	api := SlackMSG.New_SlackAPI(Cred.SLACKBOT_TOKEN)
	FB := Sheets.GetTable(Cred.GOOGLE_API_JSON, Cred.GOOGLE_SHEETS_ID, "B:E", api)
	
	PreBdays, Bdays := Sheets.Find_BDAYS(FB)

	Sheets.Send_BDAY_Private_MSG(Bdays, Cred.DB_Connect, api)
	Sheets.Prep_BDAY_MSG(PreBdays, Bdays, FB.GetIter(), api)

	return "SendDailyMSG executed without interuption"
}
