package MSGEvent

import(

)

type Request struct {
	Type	 		string `json:"Type"`
	BirthdayPerson	string `json:"BirthdayPerson"`
	SenderPerson 	string `json:"SenderPerson"`

	Message 		string `json:"Message"`

	SendPM			bool   `json:"SendPM"`
}

type Response struct{
	Message string `json:"message"`
}

func GenNotFoundResponse() Response{
	return Response{
		Message: "Implementation not found",
	}
}

func GenErrorResponse()Response{
	return Response{}
}

func GenSuccessfulDailyPingResponse() Response{
	return Response{}
}
func GenSuccessfulMessageSentResponse() Response{
	return Response{}
}