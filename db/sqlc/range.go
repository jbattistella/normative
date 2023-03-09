package db

import (
	"context"
	"math"

	"github.com/jbattistella/normative/client"
)

func GetRanges() (float64, float64, float64, error) {
	dbQueries, err := ConnectDB()

	openingPrices, err := dbQueries.GetOpeningPrice(context.Background(), "es")
	if err != nil {
		return 0, 0, 0, err
	}

	currentPrice, _, err := client.GetCurrentMarket()

	yearRange := math.Abs(openingPrices.YearOpen - currentPrice)
	monthRange := math.Abs(openingPrices.MonthOpen - currentPrice)
	weekRange := math.Abs(openingPrices.WeekOpen - currentPrice)

	return yearRange, monthRange, weekRange, nil
}
