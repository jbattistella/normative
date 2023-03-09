package db

import (
	"context"

	"github.com/jbattistella/normative/client"
)

func GetRanges() (float64, float64, float64, float64, float64, float64, error) {
	dbQueries, err := ConnectDB()

	openingPrices, err := dbQueries.GetOpeningPrice(context.Background(), "es")
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}

	currentPrice, _, err := client.GetCurrentMarket()

	yearRange := openingPrices.YearOpen - currentPrice
	monthRange := openingPrices.MonthOpen - currentPrice
	weekRange := openingPrices.WeekOpen - currentPrice

	yearChange := yearRange / openingPrices.YearOpen
	monthChange := monthRange / openingPrices.MonthOpen
	weekChange := weekRange / openingPrices.WeekOpen

	return yearRange, monthRange, weekRange, yearChange, monthChange, weekChange, nil
}
