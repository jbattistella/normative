// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: marketdays.sql

package db

import (
	"context"
	"time"
)

const getAveageRange = `-- name: GetAveageRange :one
SELECT AVG(range) 
FROM market_days
LIMIT $1
`

func (q *Queries) GetAveageRange(ctx context.Context, limit int32) (float64, error) {
	row := q.db.QueryRowContext(ctx, getAveageRange, limit)
	var avg float64
	err := row.Scan(&avg)
	return avg, err
}

const getMarketData = `-- name: GetMarketData :many
SELECT id, date, open, high, low, last, range, volume, poc3yr, poc1yr, poc0yr, poc1wk, poc1m, market_id FROM market_days
ORDER BY date ASC
`

func (q *Queries) GetMarketData(ctx context.Context) ([]MarketDay, error) {
	rows, err := q.db.QueryContext(ctx, getMarketData)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MarketDay
	for rows.Next() {
		var i MarketDay
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Open,
			&i.High,
			&i.Low,
			&i.Last,
			&i.Range,
			&i.Volume,
			&i.Poc3yr,
			&i.Poc1yr,
			&i.Poc0yr,
			&i.Poc1wk,
			&i.Poc1m,
			&i.MarketID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMarketDataByDate = `-- name: GetMarketDataByDate :many
SELECT id, date, open, high, low, last, range, volume, poc3yr, poc1yr, poc0yr, poc1wk, poc1m, market_id FROM market_days
WHERE date BETWEEN $1 AND $2
`

type GetMarketDataByDateParams struct {
	Date   time.Time `json:"date"`
	Date_2 time.Time `json:"date_2"`
}

func (q *Queries) GetMarketDataByDate(ctx context.Context, arg GetMarketDataByDateParams) ([]MarketDay, error) {
	rows, err := q.db.QueryContext(ctx, getMarketDataByDate, arg.Date, arg.Date_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MarketDay
	for rows.Next() {
		var i MarketDay
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Open,
			&i.High,
			&i.Low,
			&i.Last,
			&i.Range,
			&i.Volume,
			&i.Poc3yr,
			&i.Poc1yr,
			&i.Poc0yr,
			&i.Poc1wk,
			&i.Poc1m,
			&i.MarketID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
