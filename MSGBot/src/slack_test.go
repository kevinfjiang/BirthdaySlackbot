package main
	
import (
	"testing"
	"fmt"
    "os"
	"github.com/kevinfjiang/slackBirthdayBot/src/googleSheets"
	"github.com/kevinfjiang/slackBirthdayBot/src/slackMSG"
)


func TestSendMessage(t *testing.T){
	_, err := slackMSG.Send_MSG("Test_Message", "kfj2112@columbia.edu")
	if err != nil{
		t.Errorf("ERROR WITH SEND")
	}
}


func TestGoogleapi(t *testing.T){
	birthdayTable := googleSheets.GetTable(os.Getenv("GOOGLE_API_JSON"), os.Getenv("GOOGLE_SHEETS_ID"), "B:E")
	a,b := googleSheets.Find_BDAYS(birthdayTable)
	fmt.Println(a)
	fmt.Println(b)
}

func TestBDAYCHANNEL(t *testing.T){
	ret := slackMSG.Get_BDAY_CHANNEL()
	if ret == "" {
		t.Errorf("CHANNEL CREATION ISSUE")
	}
	fmt.Println(ret)

}