package main

import (
	"fmt"
	"github.com/kevinfjiang/slackBirthdayBot/src/SlackMSG"
	"github.com/kevinfjiang/slackBirthdayBot/src/googleSheets"
	"os"
	"testing"
)

func TestSendMessage(t *testing.T) {
	api := slackMSG.New_SlackAPI(os.Getenv("SLACKBOT_TOKEN"))
	_, err := api.Send_MSG("Test_Message", "kfj2112@columbia.edu")
	if err != nil {
		t.Errorf("ERROR WITH SEND")
	}
}

func TestGoogleapi(t *testing.T) {
	api := slackMSG.New_SlackAPI(os.Getenv("SLACKBOT_TOKEN"))
	birthdayTable := googleSheets.GetTable(os.Getenv("GOOGLE_API_JSON"), os.Getenv("GOOGLE_SHEETS_ID"), "B:E", api)
	a, b := googleSheets.Find_BDAYS(birthdayTable)
	fmt.Println(a)
	fmt.Println(b)
}

func TestBDAYCHANNEL(t *testing.T) {
	api := slackMSG.New_SlackAPI(os.Getenv("SLACKBOT_TOKEN"))
	ret := api.Get_BDAY_CHANNEL()
	if ret == "" {
		t.Errorf("CHANNEL CREATION ISSUE")
	}
	fmt.Println(ret)

}
