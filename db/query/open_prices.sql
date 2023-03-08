-- name: CreateOpeningPrice :one
INSERT INTO open_prices (
    market,
    year_open,
    month_open,
    week_open,
    Updated
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetOpeningPrice :one
SELECT * FROM open_prices
WHERE market = $1
limit 1;

-- name: UpdateYearPrice :one
UPDATE open_prices
SET year_open = $2
WHERE market = $1
RETURNING *
;

-- name: UpdateMonthPrice :one
UPDATE open_prices
SET month_open = $2
WHERE market = $1
RETURNING *
;

-- name: UpdateWeekPrice :one
UPDATE open_prices
SET week_open = $2
WHERE market = $1
RETURNING *
;







