package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	db "github.com/jbattistella/normative/db/sqlc"
	"github.com/jbattistella/normative/studies"

	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Queries *db.Queries

var DB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:pJGlBJilIdmLHvJIIFfq@containers-us-west-107.railway.app:6131/railway?sslmode=disable"
)

func main() {

	// url := "https://economiccalendar.p.rapidapi.com/events/1598072400000/1756771140000"

	// req, _ := http.NewRequest("GET", url, nil)

	// req.Header.Add("X-RapidAPI-Key", "9cd8d9ff35mshcd24dd50afc9fc4p12258ejsnd0541ec9480b")
	// req.Header.Add("X-RapidAPI-Host", "economiccalendar.p.rapidapi.com")

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer res.Body.Close()

	// var econEvents models.EconomicEvents

	// if err = json.NewDecoder(res.Body).Decode(&econEvents); err != nil {
	// 	log.Print(err)
	// }
	// for _, v := range econEvents.Events {
	// 	fmt.Println(v.Name)
	// 	fmt.Println(v.Datetime)

	// }

	var err error
	DB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	Queries = db.New(DB)

	days, _ := Queries.GetMarketData(context.Background())

	// for i := 0; i < 20; i++ {
	// 	fmt.Println(days[i].Range)
	// }

	averageRange := studies.AverageVolumeByDay(days, 20)

	for k, v := range averageRange {
		fmt.Printf("average volume for %s: %0.2f\n", k, v)
	}

	// Queries.GetMarketData(context.Background())

}
