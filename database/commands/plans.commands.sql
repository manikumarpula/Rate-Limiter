-- name: CreatePlan :one
INSERT INTO plans (
        account_id,
        name,
        rate_limit,
        burst_size,
        algorithm
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: UpdatePlan :one
UPDATE plans
SET name = $2,
    rate_limit = $3,
    burst_size = $4,
    algorithm = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;
-- name: DeletePlan :exec
DELETE FROM plans
WHERE name = $1;