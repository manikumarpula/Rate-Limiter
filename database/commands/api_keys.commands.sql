-- name: CreateAPIKey :one
INSERT INTO api_keys (
        account_id,
        name,
        description,
        key,
        status,
        created_by,
        product,
        plan_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
-- name: UpdateAPIKey :one
UPDATE api_keys
SET name = $2,
    description = $3,
    status = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;
-- name: DeleteAPIKey :exec
DELETE FROM api_keys
WHERE id = $1;