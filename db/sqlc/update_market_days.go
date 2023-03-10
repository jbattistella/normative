package db

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func UpdateMarketDays() error {
	dbQueries, err := ConnectDB()
	if err != nil {
		return err
	}

	path := "objects/onES1d(short).csv"
	csvFile, err := os.Open(path)
	if err != nil {
		return err
	}

	reader := csv.NewReader(csvFile)

	var MarketDay []CreateMarketDayParams

	var count int

	for {

		line, error := reader.Read()

		count++

		if count == 1 {
			continue
		}

		if error == io.EOF {
			break

		} else if error != nil {
			log.Fatal(error)
		}

		date, _ := time.Parse("2006/1/2", line[0])
		open, _ := strconv.ParseFloat(line[1], 64)
		last, _ := strconv.ParseFloat(line[2], 64)
		high, _ := strconv.ParseFloat(line[3], 64)
		low, _ := strconv.ParseFloat(line[4], 64)
		rng, _ := strconv.ParseFloat(line[5], 64)
		volume, _ := strconv.ParseFloat(line[6], 64)

		MarketDay = append(MarketDay, CreateMarketDayParams{
			Date:   date,
			Open:   open,
			Last:   last,
			High:   high,
			Low:    low,
			Range:  rng,
			Volume: volume,
			Market: "es",
		})

	}

	for _, v := range MarketDay {
		_, err := dbQueries.CreateMarketDay(context.Background(), v)
		if err != nil {
			return err
		}
	}
	return nil
}
