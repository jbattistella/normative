package db

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jbattistella/normative/models"
)

func TestEvents(t *testing.T) {

	jsonFile, err := os.Open("events.json")
	if err != nil {
		log.Printf("error opening file: %s", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var econEvents models.EconomicEvents

	json.Unmarshal(byteValue, &econEvents)

	for _, v := range econEvents.Events {

		layout1 := "2006/01/02"
		Date, err := time.Parse(layout1, v.Datetime[0:10])
		if err != nil {
			log.Printf("error parsing %s: %s", v.Datetime, err)
		}
		layout2 := "2006/01/02 15:04:05"
		DateTime, err := time.Parse(layout2, v.Datetime)
		if err != nil {
			log.Printf("error parsing %s: %s", v.Datetime, err)
		}

		args := CreateEventParams{
			Date:       Date,
			Time:       DateTime,
			Forecast:   v.Forecast,
			Impact:     v.Impact,
			LastUpdate: int32(v.LastUpdate),
			Name:       v.Name,
			Previous:   v.Previous,
			Region:     v.Region,
		}

		_, err = testQueries.CreateEvent(context.Background(), args)
		if err != nil {
			log.Printf("error creating event: %s", err)
		}

	}

}

func TestGetEventsWithFilter(t *testing.T) {
	args := GetEventsWithFilterParams{
		Region:   "United States",
		Impact:   "high",
		Impact_2: "medium",
	}

	res, err := testQueries.GetEventsWithFilter(context.Background(), args)
	if err != nil {
		t.Errorf("error getting events: %s", err)
	}

	if res == nil {
		t.Errorf("query response is empty: %v", res)
	}

	for _, v := range res {
		if v.Region != args.Region {
			t.Errorf("expected region %s, got %s", "United States", v.Region)
		}
		if v.Impact != args.Impact || v.Impact != args.Impact_2 {
			t.Errorf("expected impact to be %s or %s, got %s", args.Impact, args.Impact_2, v.Impact)
		}

	}
}
