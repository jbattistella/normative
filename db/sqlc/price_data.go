package db

import (
	"context"

	"github.com/jbattistella/normative/client"
)

func GetPriceData() (float64, float64, float64, float64, float64, float64, float64, error) {
	dbQueries, err := ConnectDB()

	openingPrices, err := dbQueries.GetOpeningPrice(context.Background(), "es")
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, err
	}

	currentPrice, _, err := client.GetCurrentMarket()

	yearRange := (openingPrices.YearOpen - currentPrice) * -1
	monthRange := (openingPrices.MonthOpen - currentPrice) * -1
	weekRange := (openingPrices.WeekOpen - currentPrice) * -1

	yearChange := yearRange / openingPrices.YearOpen
	monthChange := monthRange / openingPrices.MonthOpen
	weekChange := weekRange / openingPrices.WeekOpen

	return currentPrice, yearRange, monthRange, weekRange, yearChange, monthChange, weekChange, nil
}
