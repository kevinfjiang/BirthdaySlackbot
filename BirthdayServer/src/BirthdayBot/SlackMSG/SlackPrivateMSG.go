package SlackMSG

import (
	// "fmt"
	// "log"
	"sync"

	"github.com/kevinfjiang/BirthdayServer/src/DB"
	"github.com/slack-go/slack"
)

func (SA SlackAPI) Get_Private_Message(birthday map[string]interface{}, db DB.DBConnect) {
	Messages := db.Get_MSG(birthday) // Write this
	ID := birthday["SlackID"].(string)
	var wg sync.WaitGroup
	for _, message := range Messages {
		wg.Add(1)
		go func(MSG *DB.PMessage) {
			opt := genOpts(MSG) // Write thiss
			SA.PostMessage(ID, opt)
		}(message)
	}
	wg.Wait()
}

func genOpts(MSG *DB.PMessage) slack.MsgOption { 
	return nil
	// TODO set up gen oopts for every slack message
	// return slack.MsgOption{
	// 	slack.MsgOption(slack.MsgOptionText(fmt.Sprintf(msgFunc(len(birthdayPersons)), birthdayPersons...), true),),
	// 	slack.MsgOption(slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{LinkNames: len(birthdayPersons),
	// 																					IconEmoji: ":tada:",}),),
	// }
}
