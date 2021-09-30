package Sheets

import (
	"log"
	"sync"

	"github.com/kevinfjiang/BirthdayServer/src/DB"
	"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot/SlackMSG"
	"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot/fibonacci"
)

func Find_BDAYS(FB *fibHeap.FibHeap) ([]interface{}, []interface{}) {
	var prebirthday []interface{}
	var birthday []interface{}

	for _, min := FB.Minimum(); min >= 0. && min < 1.0; _, min = FB.Minimum() {
		// Checks if minimum is wihtin a day (or below 2 days) and if it is, it extracts
		// the value and adds it to a birthday/prebirthday list
		staff, _ := FB.ExtractMin()
		birthday = append(birthday, staff)
	}

	for _, min := FB.Minimum(); min >= 0. && min < 2.0; _, min = FB.Minimum() {
		staff, _ := FB.ExtractMin()
		prebirthday = append(prebirthday, staff)
	}
	log.Print("[INFO] BDAYS and PreBDAYS found")
	return prebirthday, birthday
}

func Prep_BDAY_MSG(prebirthday []interface{}, birthday []interface{}, FB *fibHeap.FibHeap, Client SlackMSG.SlackAPI) {
	if len(prebirthday) > 0 {
		prebirthdayNames := get_Bday_Names(Client, prebirthday)

		for _, bday := range append(FB.GetIter(), birthday) { // []interface{}
			Client.Send_BDAY_MSG(prebirthdayNames, bday.(*Staff).Val["SlackID"].(string), SlackMSG.Get_pre_birthdayMSG)
		}

		if len(prebirthday) > 1 {
			for i := range prebirthdayNames { 
				miniPBN := append(prebirthdayNames[:i], prebirthdayNames[i+1:]) // remove name from list!
				Client.Send_BDAY_MSG(miniPBN, prebirthday[i].(*Staff).Val["SlackID"].(string), SlackMSG.Get_pre_birthdayMSG)
			}
		}
	}

	if len(birthday) > 0 {
		birthdayNames := get_Bday_Names(Client, birthday)
		channelID := Client.Get_BDAY_CHANNEL()
		Client.Send_BDAY_MSG(birthdayNames, channelID, SlackMSG.Get_birthdayMSG)

	}
}

func get_Bday_Names(Client SlackMSG.SlackAPI, staffList []interface{}) []interface{} {
	// converts []*Staff to an []interface{} with only the names
	ret := make([]interface{}, len(staffList))

	var wg sync.WaitGroup
	for i := 0; i < len(staffList); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			linked := Client.Get_User_Link(staffList[i].(*Staff).Val["Email"].(string))
			if linked == "" && (staffList[i].(*Staff).Val["Name"] != nil) {
				linked = staffList[i].(*Staff).Val["Name"].(string)

			} else if linked == "" {
				linked = SlackMSG.Get_Anon_Name() // Add functioon to find anonNames
			}
			ret[i] = linked
		}(i)

	}
	wg.Wait()

	return ret
}

func Send_BDAY_Private_MSG(birthday []interface{}, db DB.DBConnect, Client SlackMSG.SlackAPI) {
	for _, bday := range birthday {
		Client.Get_Private_Message(bday.(*Staff).Val, db)

	}
}
