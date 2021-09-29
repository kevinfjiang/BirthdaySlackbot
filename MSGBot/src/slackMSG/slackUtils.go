package slackMSG

import (
	"fmt"
	"log"
	"sync"

	"github.com/slack-go/slack"
)


type SlackAPI struct{
	*slack.Client
}


func New_SlackAPI (Token string) SlackAPI{
	// LOG EVENT
	return SlackAPI{
		slack.New(Token),
	}
}




func (SA SlackAPI) Send_MSG(str string, email string) (string, error){
	opts := []slack.MsgOption{
		slack.MsgOption(slack.MsgOptionText(str, true),),
		slack.MsgOption(slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{LinkNames: 1,
																						IconEmoji: ":tada:",}),),
	}
				
	_, timeStamp, err := SA.PostMessage(SA.GetSlackID(email), opts...)
	return timeStamp, err
}


type msgFunction func(i int)(string)

func (SA SlackAPI) Send_BDAY_MSG(birthdayPersons []interface{}, ID string, msgFunc msgFunction)(string, error){
	opts := []slack.MsgOption{
		slack.MsgOption(slack.MsgOptionText(fmt.Sprintf(msgFunc(len(birthdayPersons)), birthdayPersons...), true),),
		slack.MsgOption(slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{LinkNames: len(birthdayPersons), 
																						IconEmoji: ":tada:",}),),
	}

	_, timeStamp, err := SA.PostMessage(ID, opts...)   //Don't forget the time stamp, use that
	if err == nil{
		return timeStamp, nil 
	}
	return timeStamp, err
}

func (SA SlackAPI) Get_BDAY_CHANNEL() string {
	// LOG EVENTT
	channelList, _, err  := SA.GetConversations(&slack.GetConversationsParameters{ExcludeArchived: true})
	
	if err == nil{
		for _, channel := range(channelList){
			if channel.Name == "birthday"{
				return channel.ID
			}
		}
	}

	channel, err2 := SA.CreateConversation("birthday", false)
	
	if err2!=nil{
		// LOG EVENT
		return ""
	}

	return channel.ID
}

func (SA SlackAPI) GetSlackID(Email string) string {
	// LOG SEARCH
	userID, err := SA.GetUserByEmail(Email)
	if err != nil{
		// LOG EVENT
		return "No ID Found"
	}
	return userID.ID
}

func (SA SlackAPI) Get_User_Link(Email string) string {
	userID, err := SA.GetUserByEmail(Email)
	if err != nil{
		return ""
	}
	return "@" + userID.Name

}
