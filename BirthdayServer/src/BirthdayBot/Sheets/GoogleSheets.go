package Sheets

import (
	"fmt"
	"log"

	"math"
	"strings"
	"sync"
	"time"

	"context"
	"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot/SlackMSG"
	"github.com/kevinfjiang/BirthdayServer/src/BirthdayBot/fibonacci"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Staff struct {
	Val map[string]interface{}
}

func (S *Staff) Tag() interface{} {
	return S
}

var TodayTime = time.Now().Truncate(24 * time.Hour) // Repeatedly calculated, gets current day but starts at midnight

func (S *Staff) Key() float64 {
	// returns how far away the users next bday is to the current day
	bdaystring := S.Val["Birthday"].(string)

	if bdaystring[:len(bdaystring)-5] == "2/29" { // Gotta love leap years
		bdaystring = "2/28/0000"
	}

	YEAR := fmt.Sprint(TodayTime.Year())

YearChange: // Only comes if the date has already passed
	BDAY, err := time.Parse("1/2/2006", bdaystring[:len(bdaystring)-4]+YEAR)

	if err != nil || S.Val["OthersNotify"] == "No" {
		return math.MaxFloat64 // Fibonnaci heap, places at way back
	}

	distance := BDAY.Sub(TodayTime).Hours() / 24

	if distance < 0 && fmt.Sprint(TodayTime.Year()) == YEAR { // if date passed or more than one iteration
		YEAR = fmt.Sprint(TodayTime.AddDate(1, 0, 0).Year())
		goto YearChange
	}

	return distance
}

func GetService(jsonPath string) *sheets.Service {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(jsonPath))

	if err != nil {
		log.Fatalln(err)
	}
	return srv
}

func GetTable(jsonPath string, spreadsheetId string, readRange string, Client SlackMSG.SlackAPI) *fibHeap.FibHeap {
	log.Printf("[INFO] Starting connectiono to GoogleSheets DB")
	srv := GetService(jsonPath)
	template := strings.Replace(readRange, ":", "%s:", 1) + "%s" //Format of a query for a column

	colNames := fmt.Sprintf(template, []interface{}{"1", "1"}...)
	tableContent := fmt.Sprintf(template, []interface{}{"2", ""}...) // TODO BUGSS HERE FIGURE THIS OOUT Find a way to query only necessary rows

	rows, err1 := srv.Spreadsheets.Values.Get(spreadsheetId, colNames).Do()
	table, err2 := srv.Spreadsheets.Values.Get(spreadsheetId, tableContent).DateTimeRenderOption("SERIAL_NUMBER").Do()
	// LOG EVENTs
	if err1 != nil || err2 != nil {
		log.Fatalln("[ERROR] Unable to retrieve data from sheet")
	}

	if len(table.Values) == 0 {
		log.Print("[Warning] Table exists but table is empty")
		return nil

	} else {
		log.Print("[INFO] Table Found, organizing table")
		StaffTable := fibHeap.NewFibHeap() // FibHeaps, ordered and log(n) extract times, only 2, iteration through time is the same as slices
		var wg sync.WaitGroup
		// LOG EVENT
		for i, row := range table.Values {
			staffMap := make(map[string]interface{})

			for ii, colName := range rows.Values[0] { // Uses column headers to assign the map
				staffMap[colName.(string)] = row[ii]
			}

			if ID := staffMap["SlackID"]; ID == nil || ID == "" || ID == "No ID Found" {
				wg.Add(1)
				go func(index int, staffrow []interface{}, email string) {
					defer wg.Done()
					id := Client.GetSlackID(email)
					staffMap["SackID"] = id
					log.Println("[INFO] %s got new id %s in row %d", email, id, index)
					write(srv, spreadsheetId, append(staffrow, id), index)
				}(i+2, row, staffMap["Email"].(string)) // Rows start at one and header so +2
			}
			StaffTable.InsertValue(&Staff{staffMap})

		}
		wg.Wait()

		log.Print("[INFO] Google sheets table read into fibHeap")
		return StaffTable
	}
}

func write(srv *sheets.Service, spreadsheetId string, staffRow []interface{}, rowNumber int) {
	var vr sheets.ValueRange
	writeRange := fmt.Sprint("B", rowNumber, ":", rowNumber)

	vr.Values = append(vr.Values, staffRow)

	_, err := srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalln("Unable to retrieve data from sheet. %v", err)
	}
}
