package MSGEvent

import(
	"encoding/json"
)

type Request struct {
	Resource              string                        `json:"resource"` // The resource path defined in API Gateway
    Path                  string                        `json:"path"`     // The url path for the caller
    HTTPMethod            string                        `json:"httpMethod"`
    Headers               map[string]string             `json:"headers"`
    PathParameters        map[string]string             `json:"pathParameters"`
    StageVariables        map[string]string             `json:"stageVariables"`
    Body                  string                        `json:"body"`
    IsBase64Encoded       bool                          `json:"isBase64Encoded,omitempty"`


	Type	 		string `json:"Type"`
	BirthdayPerson	string `json:"BirthdayPerson"`
	SenderPerson 	string `json:"SenderPerson"`

	Message 		string `json:"Message"`

	SendPM			bool   `json:"SendPM"`
}

type Response struct{
	Status 			int    			  `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
    Body            []byte            `json:"body"`
    IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
}

func GenBody(responses interface{})([]byte){
	if messsage, err := json.Marshal(responses); err==nil{
		return messsage
	}
	return []byte("Error while parsing map")

}

func GenNotFoundResponse() Response{
	return Response{
		Status:	 404, 
		Headers: map[string]string {"Content-Type": "application/json",},
		Body: GenBody(map[string]string{"message": "Web page not found"}),
		IsBase64Encoded: true,
	}
}

func GenErrorResponse()Response{
	return Response{
		Status:	 418, 
		Headers: map[string]string {"Content-Type": "application/json",},
		Body: GenBody(map[string]string{"message": "Implementation not found"}),
		IsBase64Encoded: true,
	}
}

func GenSuccessfulDailyPingResponse() Response{
	return Response{}
}
func GenSuccessfulMessageSentResponse() Response{
	return Response{}
}