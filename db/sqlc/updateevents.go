package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jbattistella/normative/client"

	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func UpdateEvents() error {

	DB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer DB.Close()

	dbQueries = New(DB)

	econEvents := client.GetEvents()

	LastDate, err := dbQueries.GetLastEventDate(context.Background())
	if err != nil {
		return err
	}

	for _, v := range econEvents.Events {
		layout1 := "2006/01/02"
		Date, err := time.Parse(layout1, v.Datetime[0:10])
		if err != nil {
			log.Printf("error parsing %s: %s", v.Datetime, err)
		}
		layout2 := "2006/01/02 15:04:05"
		dateTime, err := time.Parse(layout2, v.Datetime)
		if err != nil {
			log.Printf("error parsing %s: %s", v.Datetime, err)
		}

		if LastDate.Before(Date) {

			ep := CreateEventParams{
				Date:       Date,
				Time:       dateTime,
				Forecast:   v.Forecast,
				Impact:     v.Impact,
				LastUpdate: int32(v.LastUpdate),
				Name:       v.Name,
				Previous:   v.Previous,
				Region:     v.Region,
			}

			_, err = dbQueries.CreateEvent(context.Background(), ep)

			if err != nil {
				return err
			}

		}

	}
	return nil
}
