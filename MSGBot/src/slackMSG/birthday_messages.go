package slackMSG

import (
    "math/rand"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano()) // random selects
}

var BDAY_MESSAGE = map[int][]string{
	1: {"Hello %s %s", "Hello2%s %s", "Hello3%s %s", "4%s %s", "5%s %s"},
	2: {"Hello %s %s", "Hello2%s %s", "Hello3%s %s", "4%s %s", "5%s %s"},
}

var PRE_BDAY_MESSAGE = map[int][]string{
	1: {"Hello %s %s", "Hello2%s %s", "Hello3%s %s", "4%s %s", "5%s %s"},
	2: {"Hello %s %s", "Hello2%s %s", "Hello3%s %s", "4%s %s", "5%s %s"},
}

var AnonName = []string{"A special someone", "An oldie fella", "An oldie but a goodie", "Gramps", "Special person"}



func Get_birthdayMSG(BDAYPEOPLE int) (string){
	return BDAY_MESSAGE[BDAYPEOPLE][rand.Intn(len(BDAY_MESSAGE[BDAYPEOPLE]))]
}

func Get_pre_birthdayMSG(countVIPS int) (string){
	return PRE_BDAY_MESSAGE[countVIPS][rand.Intn(len(PRE_BDAY_MESSAGE[countVIPS]))]
}

func Get_Anon_Name()string{
	return AnonName[rand.Intn(len(AnonName))]
}