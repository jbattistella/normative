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

	ev, err := ep.GetEvents()
	if err != nil {
		log.Println(err)
	}

	for _, v := range ev {
		date := v.Date.String()[0:10]
		time := v.Time.String()[12:19]
		fmt.Printf("%s | %s | %s | %s \n", date, time, v.Impact, v.Name)
	}

}
