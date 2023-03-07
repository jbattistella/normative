-- name: CreateEvent :one
INSERT INTO events (
    date,
    time,
    forecast,
    impact,
    last_update,
    name,
    previous,
    region
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetEventByDate :many
SELECT * FROM events
WHERE date = $1;

-- name: GetLastEventDate :one
SELECT date FROM events
ORDER BY date DESC LIMIT 1;

-- name: GetEventsWithFilter :many
SELECT * FROM events
WHERE 
    region = $1
AND
    impact = $2
AND
    impact = $3;




