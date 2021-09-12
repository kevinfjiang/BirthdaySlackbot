package googleSheets


import(
	"sync"
	"github.com/kevinfjiang/slackBirthdayBot/src/fibonacci"
	"github.com/kevinfjiang/slackBirthdayBot/src/slackMSG"
)

func Find_BDAYS(FB *fibHeap.FibHeap) ([]interface{}, []interface{}){
	prebirthday := []interface{}{}
	birthday := []interface{}{}

	for _, min := FB.Minimum(); min >= 0. && min < 1.0; _, min = FB.Minimum(){
		// Checks if minimum is wihtin a day (or below 2 days) and if it is, it extracts
		// the value and adds it to a birthday/prebirthday list
		staff, _ := FB.ExtractMin()
		birthday = append(birthday, staff)
	} 

	for _, min := FB.Minimum(); min >= 0. && min < 2.0; _, min = FB.Minimum(){
		staff, _ := FB.ExtractMin()
		prebirthday = append(prebirthday, staff)
	}

	return prebirthday, birthday
}


func get_Bday_Names(staffList []interface{}) []interface{}{
	// converts []*Staff to an []interface{} with only the names
	ret := make([]interface{}, len(staffList))
	
	var wg sync.WaitGroup
    for i:=0; i<len(staffList); i++ {
		wg.Add(1)

		go func(i int){
			defer wg.Done()
			linked := slackMSG.Get_User_Link(staffList[i].(*Staff).Val["Email"].(string))
			if linked == "" && (staffList[i].(*Staff).Val["Name"] != nil){
				linked = staffList[i].(*Staff).Val["Name"].(string)
			} else if linked == ""{
				linked = slackMSG.Get_Anon_Name() // Add functioon to find anonNames
			}
			ret[i] = linked
		}(i)
		
    }
	wg.Wait()
	
	return ret
}
func Prep_BDAY_MSG(prebirthday []interface{}, birthday []interface{}, FB *fibHeap.FibHeap) {
	if len(prebirthday)>0{
		prebirthdayNames := get_Bday_Names(prebirthday) 

		for _, stff := range(append(FB.GetIter(), birthday)){//slice of interface{}
			slackMSG.Send_BDAY_MSG(prebirthdayNames, stff.(*Staff).Val["SlackID"].(string), slackMSG.Get_pre_birthdayMSG)
		}

		if len(prebirthday) > 1{
			for i := range(prebirthdayNames){ // remove name from list!
				miniPBN := append(prebirthdayNames[:i], prebirthdayNames[i+1:])
				slackMSG.Send_BDAY_MSG(miniPBN, prebirthday[i].(*Staff).Val["SlackID"].(string), slackMSG.Get_pre_birthdayMSG)	
			} 
		}
	}

	if len(birthday) > 0{
		birthdayNames := get_Bday_Names(birthday)
		channelID := slackMSG.Get_BDAY_CHANNEL()
		slackMSG.Send_BDAY_MSG(birthdayNames, channelID, slackMSG.Get_birthdayMSG)

	}
}