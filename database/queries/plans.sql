-- name: GetPlanByName :one
SELECT *
FROM plans
WHERE name = $1;