package slackMSG

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
)


type SlackAPI struct{
	client *slack.Client
}


func New_SlackAPI (Token string) *SlackAPI{
	// LOG EVENT
	return &SlackAPI{
		slack.New(Token),
	}
}


type msgFunction func(i int)(string)

func (SA *SlackAPI) Send_BDAY_MSG(birthdayPersons []interface{}, ID string, msgFunc msgFunction)(string, error){
	opts := []slack.MsgOption{slack.MsgOption(
			slack.MsgOptionText(fmt.Sprintf(msgFunc(len(birthdayPersons)), birthdayPersons...), true),),
		slack.MsgOption(
			slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{LinkNames: len(birthdayPersons),
			IconEmoji: ":tada:",}),)}

	_,timeStamp, err := SA.client.PostMessage(ID, opts...)   //Don't forget the time stamp, use that
	if err == nil{
		return timeStamp, nil 
	}
	return timeStamp, err
}

func (SA *SlackAPI) Get_BDAY_CHANNEL() string {
	//LOG EVENTT
	channelList, _, err  := SA.client.GetConversations(&slack.GetConversationsParameters{ExcludeArchived: true})
	
	if err == nil{
		for _,channel := range(channelList){
			if channel.Name == "birthday"{
				return channel.ID
			}
		}
	}

	channel, err2 := SA.client.CreateConversation("birthday", false)
	
	if err2!=nil{
		// LOG EVENT
		return ""
	}

	return channel.ID
}

func (SA *SlackAPI) Send_MSG(str string, email string)(string, error){
	opts := []slack.MsgOption{slack.MsgOption(
						slack.MsgOptionText(str, true),),
			  slack.MsgOption(
						slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{LinkNames: 1,
						IconEmoji: ":tada:",}),)}
				
	_,timeStamp, err := SA.client.PostMessage(SA.GetSlackID(email), opts...)
	return timeStamp, err
}

func (SA *SlackAPI) GetSlackID(Email string) string {
	// LOG SEARCH
	userID, err := SA.client.GetUserByEmail(Email)
	if err != nil{
		// LOG EVENT
		return "No ID Found"
	}
	return userID.ID
}

func (SA *SlackAPI) Get_User_Link(Email string) string {
	userID, err := SA.client.GetUserByEmail(Email)
	if err != nil{
		return ""
	}
	return "@" + userID.Name

}