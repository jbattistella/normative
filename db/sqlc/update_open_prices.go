package db

import (
	"context"
	"fmt"
	"log"
	"time"
)

func UpdateOpenPrices() error {
	dbQueries, err := ConnectDB()
	if err != nil {
		return err
	}

	//if the market is open today then ..
	last, err := dbQueries.GetLastMarketRecord(context.Background())

	openPrices, err := dbQueries.GetOpeningPrice(context.Background(), "es")

	//Get Year Open
	year := time.Now().Year()

	if last.Date.Year() != openPrices.Updated.Year() {
		fmt.Println("year")

		date, err := time.Parse("2006-01-02", fmt.Sprintf("%d-01-01", year))
		if err != nil {
			log.Println(err)
		}
		date2, err := time.Parse("2006-01-02", fmt.Sprintf("%d-01-05", year))
		if err != nil {
			log.Println(err)
		}

		fmt.Println(date, date2)

		args := GetMarketDataByDateRangeParams{
			Date:   date,
			Date_2: date2,
		}

		marketDay, err := dbQueries.GetMarketDataByDateRange(context.Background(), args)

		if err != nil {
			log.Println(err)
		}

		fmt.Println(marketDay)

		args2 := UpdateYearPriceParams{Market: "es", YearOpen: marketDay[0].Open}

		_, err = dbQueries.UpdateYearPrice(context.Background(), args2)
		if err != nil {
			return err
		}

	}

	if last.Date.Month() != openPrices.Updated.Month() {
		fmt.Println("month")
		year, month, _ := time.Now().Date()
		var format string
		if month < 10 {
			format = "2006-1-02"
		} else {
			format = "2006-01-02"
		}
		date, err := time.Parse(format, fmt.Sprintf("%d-%d-01", year, month))
		if err != nil {
			log.Println(err)
		}
		date2, err := time.Parse(format, fmt.Sprintf("%d-%d-05", year, month))
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

		args2 := UpdateMonthPriceParams{Market: "es", MonthOpen: marketDay[0].Open}

		_, err = dbQueries.UpdateMonthPrice(context.Background(), args2)
		if err != nil {
			return err
		}
	}

	_, lastMarketDayLog := last.Date.ISOWeek()
	_, lastOpenPriceUpdate := openPrices.Updated.ISOWeek()

	if lastOpenPriceUpdate != lastMarketDayLog {
		fmt.Println("week")
		var thisWeek []MarketDay

		marketDays, err := dbQueries.GetMarketDataByDays(context.Background(), 7)
		if err != nil {
			return err
		}

		_, currentWeek := time.Now().ISOWeek()

		fmt.Println(marketDays)

		for _, v := range marketDays {
			_, w := v.Date.ISOWeek()
			if w == currentWeek {
				thisWeek = append(thisWeek, v)
			}
		}
		fmt.Println(thisWeek)

		weekOpen := thisWeek[0].Open

		args2 := UpdateWeekPriceParams{Market: "es", WeekOpen: weekOpen}

		_, err = dbQueries.UpdateWeekPrice(context.Background(), args2)
		if err != nil {
			return err
		}
	}

	return err

}
