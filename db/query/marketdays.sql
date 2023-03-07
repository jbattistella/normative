-- name: GetMarketDataByDateRange :many
SELECT * FROM market_days
WHERE date BETWEEN $1 AND $2;

-- name: GetMarketDataByDate :one
SELECT * FROM market_days
WHERE date = $1;

-- name: GetMarketDataByDays :many
SELECT * FROM market_days
ORDER BY date ASC
LIMIT $1;

-- name: GetLastMarketRecord :one
SELECT * FROM market_days
ORDER BY date ASC
LIMIT 1;

-- name: GetAveageRange :one
SELECT AVG(range) 
FROM market_days
LIMIT $1;