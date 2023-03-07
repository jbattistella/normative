package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func updateOpenPrices() error {
	DB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	dbQueries = New(DB)

	//if the market is open today then ..
	last, err := dbQueries.GetLastMarketRecord(context.Background())

	openPrices, err := dbQueries.GetOpeningPrice(context.Background(), "ES")

	//Get Year Open
	year := time.Now().Year()

	if last.Date.Year() != openPrices.Updated.Year() {

		date, err := time.Parse("2006-01-02", fmt.Sprintf("%d-01-01", year))
		if err != nil {
			log.Println(err)
		}
		date2, err := time.Parse("2006-01-02", fmt.Sprintf("%d-01-05", year))
		if err != nil {
			log.Println(err)
		}

		args := GetMarketDataByDateRangeParams{
			Date:   date,
			Date_2: date2,
		}

		marketDay, err := dbQueries.GetMarketDataByDateRange(context.Background(), args)

		if err != nil {
			log.Println(err)
		}

		args2 := UpdateYearPriceParams{Market: "ES", YearOpen: marketDay[0].Open}

		_, err = dbQueries.UpdateYearPrice(context.Background(), args2)
		if err != nil {
			return err
		}

	}

	if last.Date.Month() != openPrices.Updated.Month() {
		year, month, _ := time.Now().Date()
		date, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%d-01", year, month))
		if err != nil {
			log.Println(err)
		}
		date2, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%d-05", year, month))
		if err != nil {
			log.Println(err)
		}

		args := GetMarketDataByDateRangeParams{
			Date:   date,
			Date_2: date2,
		}

		marketDay, err := dbQueries.GetMarketDataByDateRange(context.Background(), args)

		if err != nil {
			log.Println(err)
		}

		args2 := UpdateMonthPriceParams{Market: "ES", MonthOpen: marketDay[0].Open}

		_, err = dbQueries.UpdateMonthPrice(context.Background(), args2)
		if err != nil {
			return err
		}
	}

	_, week := last.Date.ISOWeek()
	_, currentWeek := openPrices.Updated.ISOWeek()

	if currentWeek != week {

		var thisWeek []MarketDay

		marketDays, err := dbQueries.GetMarketDataByDays(context.Background(), 7)
		if err != nil {
			return err
		}

		for _, v := range marketDays {
			_, w := v.Date.ISOWeek()
			if w == currentWeek {
				thisWeek = append(thisWeek, v)
			}
		}

		weekOpen := thisWeek[0].Open

		args2 := UpdateWeekPriceParams{Market: "ES", WeekOpen: weekOpen}

		_, err = dbQueries.UpdateWeekPrice(context.Background(), args2)
		if err != nil {
			return err
		}
	}

	return err

}
