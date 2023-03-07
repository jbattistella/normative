package main

import (
	"flag"
	"fmt"
	"log"

	db "github.com/jbattistella/normative/db/sqlc"
	"github.com/jbattistella/normative/engine"
)

var Region string
var Impact string
var Update bool

func init() {

	flag.BoolVar(&Update, "u", false, "Update the database")

	flag.StringVar(&Region, "Region", "United States", "a string var")

	flag.StringVar(&Impact, "Impact", "high", "a string var")

	flag.Parse()

}

func main() {

	//todo: create flags to update db, pass in args for event filter, and choose study's to display
	//todo: get quote filter

	if Update {
		err := db.UpdateEvents()
		if err != nil {
			log.Println(err)
		}
	}

	ep := engine.NewEventParams()

	ep.Region = Region
	ep.Impact = append(ep.Impact, Impact)

	// err := db.UpdateEvents()
	// if err != nil {
	// 	log.Println(err)
	// }

	ev, err := ep.GetEvents()
	if err != nil {
		log.Println(err)
	}

	for _, v := range ev {
		date := v.Date.String()[0:10]
		time := v.Time.String()[12:19]
		fmt.Printf("%s | %s | %s | %s \n", date, time, v.Impact, v.Name)
	}

	marketDays, err := engine.GetMarketByDate("2022-12-31", "2023-01-03")

	for _, v := range marketDays {
		fmt.Println(v.Open)
	}

	// ee := client.GetEvents()

	// fmt.Println(ee.Events)

	// Queries.GetMarketData(context.Background())

}
