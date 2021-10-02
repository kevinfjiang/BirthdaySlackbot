package SlackMSG

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
)

type SlackAPI struct {
	*slack.Client
}

func New_SlackAPI(Token string) SlackAPI {
	log.Print("[INFO] Connecting to SlackAPI with token")
	return SlackAPI{
		slack.New(Token),
	}
}

type msgFunction func(i int) string

func (SA SlackAPI) Send_BDAY_MSG(birthdayPersons []interface{}, ID string, msgFunc msgFunction) (string, error) {
	opts := []slack.MsgOption{
		slack.MsgOption(slack.MsgOptionText(fmt.Sprintf(msgFunc(len(birthdayPersons)), birthdayPersons...), true)),
		slack.MsgOption(slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{LinkNames: len(birthdayPersons),
			IconEmoji: ":tada:"})),
	}

	_, timeStamp, err := SA.PostMessage(ID, opts...) //Don't forget the time stamp, use that
	if err == nil {
		log.Printf("[WARNING] Error sending message to this ID %s", ID)
		return timeStamp, nil
	}
	
	return timeStamp, err
}

func (SA SlackAPI) Get_BDAY_CHANNEL() string {
	log.Print("[INFO] Finding channel named 'birthday'")
	channelList, _, err := SA.GetConversations(&slack.GetConversationsParameters{ExcludeArchived: true})

	if err == nil {
		for _, channel := range channelList {
			if channel.Name == "birthday" {
				log.Print("[INFO] Channel Found!")
				return channel.ID
			}
		}
	}

	channel, err2 := SA.CreateConversation("birthday", false)

	if err2 != nil {
		log.Fatalln("[ERROR] No birthday channel found, Create one or check credentials")
	}

	return channel.ID
}

func (SA SlackAPI) GetSlackID(Email string) string {
	userID, err := SA.GetUserByEmail(Email)
	if err != nil {
		log.Printf("[INFO] No ID Found for %s", Email)
		return "No ID Found"
	}
	return userID.ID
}

func (SA SlackAPI) Get_User_Link(IdentifierString string, Type string) string {
	var userID *slack.User
	var err error = nil
	if Type == "Email"{
		userID, err = SA.GetUserByEmail(IdentifierString)
	} else if Type == "ID" {
		userID, err = SA.GetUserInfo(IdentifierString)
	}
	if err != nil || userID == nil{
		log.Printf("[INFO] Username not found for %s", IdentifierString)
		return ""
	}
	return "@" + userID.Name

}
