package googleSheets

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"math"
	"sync"

	"github.com/kevinfjiang/slackBirthdayBot/src/slackMSG"
	"github.com/kevinfjiang/slackBirthdayBot/src/fibonacci"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Staff struct{
	Val 	map[string]interface{}
}

func (S *Staff) Tag() interface{} {
	return S
}

var TodayTime = time.Now().Truncate(24 * time.Hour) // Repeatedly calculated, gets current day but starts at midnight

func (S *Staff) Key() float64 { 
	// returns how far away the users next bday is to the current day
	bdaystring := S.Val["Birthday"].(string)

	if bdaystring[:len(bdaystring)-5] == "2/29"{ // Gotta love leap years
		bdaystring ="2/28/0000"
	}

	YEAR:=fmt.Sprint(TodayTime.Year())

	YearChange: // Only comes if the date has already passed
		BDAY, err := time.Parse("1/2/2006", bdaystring[:len(bdaystring)-4]+YEAR)

		if err!=nil || S.Val["OthersNotify"]=="No"{
			return math.MaxFloat64 // Fibonnaci heap, places at way back
		}

		distance:=BDAY.Sub(TodayTime).Hours()/24
		
		if distance < 0 && fmt.Sprint(TodayTime.Year())==YEAR{ // if date passed or more than one iteration
			YEAR = fmt.Sprint(TodayTime.AddDate(1,0,0).Year())
			goto YearChange
		}

	return distance
}


func GetService(jsonPath string)(*sheets.Service){
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(jsonPath))
	
	if err != nil {
		log.Fatal(err)
	}
	return srv
}


func GetTable(jsonPath string, spreadsheetId string, readRange string, Client *slackMSG.SlackAPI) (*fibHeap.FibHeap) {
	srv := GetService(jsonPath)
	template := strings.Replace(readRange, ":", "%s:", 1) + "%s" //Format of a query for a column
	
	colNames := fmt.Sprintf(template, []interface{}{"1","1"}...)
	tableContent := fmt.Sprintf(template, []interface{}{"2", ""}...)

	collumns, err1 :=  srv.Spreadsheets.Values.Get(spreadsheetId, colNames).Do()
	table, err2 := srv.Spreadsheets.Values.Get(spreadsheetId, tableContent).DateTimeRenderOption("SERIAL_NUMBER").Do()
	
	if err1 != nil || err2 != nil {
		log.Fatalf("Unable to retrieve data from sheet")
	}

	if len(table.Values) == 0 {
		return nil
	
	} else {
		StaffTable := fibHeap.NewFibHeap() // FibHeaps, ordered and log(n) extract times, only 2, iteration through time is the same as slices
		var wg sync.WaitGroup

		for _, col := range(table.Values){
			staffMap := make(map[string]interface{})

			for i, colName := range(collumns.Values[0]){ // Uses column headers to assign the map
				staffMap[colName.(string)] = col[i]
			}

			if staffMap["SlackID"] == nil{
				wg.Add(1)
				go func(){
					defer wg.Done()
					staffMap["SlackID"] = Client.GetSlackID(staffMap["Email"].(string))
				}()
			}
			StaffTable.InsertValue(&Staff{staffMap})
			
		}
		wg.Wait()

		return StaffTable
	}
}