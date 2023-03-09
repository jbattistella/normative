package db

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

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

func (e *EventParams) GetEvents() ([]Event, error) {

	queries, err := ConnectDB()
	if err != nil {
		return []Event{}, err
	}

	today := time.Now()

	events, _ := queries.GetEventByDate(context.Background(), today)

	var todayEvents []Event

	for _, v := range events {

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
