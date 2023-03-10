package client

import (
	"github.com/piquette/finance-go/quote"
)

func GetCurrentMarket() (float64, float64, error) {
	q, err := quote.Get("ES=F")
	if err != nil {
		return 0, 0, err
	}

	return q.RegularMarketPrice, q.RegularMarketChangePercent, nil
}
