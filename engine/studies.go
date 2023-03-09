package engine

import (
	db "github.com/jbattistella/normative/db/sqlc"
)

func MovingAverage(values []db.MarketDay, lookback int) float64 {
	var count float64
	var total float64
	for i := 0; i < lookback; i++ {
		count++
		total = total + float64(values[i].Range)
	}

	average := total / count

	return average
}

// func AverageVolumeByDay(values []db.MarketDay, lookback int) map[string]float32 {
// 	volumeMap := make(map[string]float32)

// 	weekdayCounter := make(map[string]float32)

// 	for i := 0; i < lookback; i++ {
// 		wk := values[i].Date.Weekday().String()
// 		volumeMap[wk] += values[i].Volume
// 		weekdayCounter[wk]++
// 	}

// 	for k, _ := range volumeMap {
// 		volumeMap[k] = volumeMap[k] / weekdayCounter[k]
// 	}

// 	return volumeMap

// }
