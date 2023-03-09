package engine

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	db "github.com/jbattistella/normative/db/sqlc"
	"github.com/jbattistella/normative/models"

	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type EventParams struct {
	Region string
	Impact []string
}

func NewEventParams() *EventParams {
	ep := &EventParams{}
	return ep

}

func getEventsJson() models.EconomicEvents {
	jsonFile, err := os.Open("events.json")
	if err != nil {
		log.Printf("error opening file: %s", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var econEvents models.EconomicEvents

	json.Unmarshal(byteValue, &econEvents)
	return econEvents
}

func (e *EventParams) GetEvents() ([]db.Event, error) {

	queries, err := db.ConnectDB()
	if err != nil {
		return []db.Event{}, err
	}

	today := time.Now()

	events, _ := queries.GetEventByDate(context.Background(), today)

	// eEvents := getEventsJson()

	var todayEvents []db.Event

	for _, v := range events {
		// layout2 := "2006/01/02 15:04:05"
		// DateTime, err := time.Parse(layout2, v.Datetime)
		// if err != nil {
		// 	log.Printf("error parsing %s: %s", v.Datetime, err)
		// }

		// y1, m1, d1 := DateTime.Date()
		// y2, m2, d2 := time.Now().Date()

		// eventDate := fmt.Sprintf("%d, %v, %d", y1, m1, d1)
		// today := fmt.Sprintf("%d, %v, %d", y2, m2, d2)

		// fmt.Println(DateTime, time.Now())

		// if eventDate == today {
		// 	fmt.Println(DateTime, time.Now())

		switch {

		case len(e.Impact) == 2:
			if v.Impact == e.Impact[0] || v.Impact == e.Impact[1] && v.Region == e.Region {
				todayEvents = append(todayEvents, v)
			}

		case len(e.Impact) == 1:
			if v.Impact == e.Impact[0] && v.Region == e.Region {
				todayEvents = append(todayEvents, v)
			}

		default:
			if v.Region == e.Region {
				todayEvents = append(todayEvents, v)
			}

		}

	}

	return todayEvents, nil
}

func GetMarketByDateRange(d string, d2 string) ([]db.MarketDay, error) {

	queries, err := db.ConnectDB()
	if err != nil {
		return []db.MarketDay{}, err
	}

	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		log.Println(err)
	}
	date2, err := time.Parse("2006-01-02", d2)
	if err != nil {
		log.Println(err)
	}

	args := db.GetMarketDataByDateRangeParams{
		Date:   date,
		Date_2: date2,
	}

	marketDay, err := queries.GetMarketDataByDateRange(context.Background(), args)

	if err != nil {
		log.Println(err)
	}

	return marketDay, nil

}
