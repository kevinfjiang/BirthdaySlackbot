package slackMSG

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)



var Token string = os.Getenv("SLACKBOT_TOKEN") // set as os variable
var slackAPI *slack.Client = slack.New(Token)


type msgFunction func(i int)(string)

func Send_BDAY_MSG(birthdayPersons []interface{}, ID string, msgFunc msgFunction)(string, error){
	opts := []slack.MsgOption{slack.MsgOption(
			slack.MsgOptionText(fmt.Sprintf(msgFunc(len(birthdayPersons)), birthdayPersons...), true),),
		slack.MsgOption(
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{LinkNames: len(birthdayPersons),
			IconEmoji: ":tada:",}),)}

	_,timeStamp, err := slackAPI.PostMessage(ID, opts...)   //Don't forget the time stamp, use that
	if err == nil{
		return timeStamp, nil 
	}
	return timeStamp, err
}

func Get_BDAY_CHANNEL() string {

	channelList, _, err  := slackAPI.GetConversations(&slack.GetConversationsParameters{ExcludeArchived: true})
	
	if err == nil{
		for _,channel := range(channelList){
			if channel.Name == "birthday"{
				return channel.ID
			}
		}
	}

	channel, err2 := slackAPI.CreateConversation("birthday", false)
	
	if err2!=nil{
		fmt.Println(err2)
		return ""
	}

	return channel.ID
}

func Send_MSG(str string, email string)(string, error){
	opts := []slack.MsgOption{slack.MsgOption(
						slack.MsgOptionText(str, true),),
			  slack.MsgOption(
						slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{LinkNames: 1,
						IconEmoji: ":tada:",}),)}
				
	_,timeStamp, err := slackAPI.PostMessage(GetSlackID(email), opts...)
	return timeStamp, err
}

func GetSlackID(Email string) string {
	userID, err := slackAPI.GetUserByEmail(Email)
	if err != nil{
		return ""
	}
	return userID.ID
}

func Get_User_Link(Email string) string {
	userID, err := slackAPI.GetUserByEmail(Email)
	if err != nil{
		return ""
	}
	return "@" + userID.Name

}