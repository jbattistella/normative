-- name: CreateMarketDay :one
INSERT INTO market_days (
    date,
    open,
    high,
    low,
    last,
    range,
    volume,
    market
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;


-- name: GetMarketDataByDateRange :many
SELECT * FROM market_days
WHERE date BETWEEN $1 AND $2;

-- name: GetMarketDataByDate :one
SELECT * FROM market_days
WHERE date = $1;

-- name: GetMarketDataByDays :many
SELECT * FROM market_days
ORDER BY date DESC
LIMIT $1;

-- name: GetLastMarketRecord :one
SELECT * FROM market_days
ORDER BY date ASC
LIMIT 1;

-- name: GetAveageRange :one
SELECT AVG(range) 
FROM market_days
LIMIT $1;