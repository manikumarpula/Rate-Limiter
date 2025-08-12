-- name: GetAPIKey :one
SELECT *
FROM api_keys
WHERE key = $1;