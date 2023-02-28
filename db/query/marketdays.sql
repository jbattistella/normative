-- name: GetMarketDataByDate :many
SELECT * FROM market_days
WHERE date BETWEEN $1 AND $2;

-- name: GetMarketData :many
SELECT * FROM market_days
ORDER BY date ASC;

-- name: GetAveageRange :one
SELECT AVG(range) 
FROM market_days
LIMIT $1;