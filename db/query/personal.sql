-- name: CreatePersonal :one
INSERT INTO personal (
  account_id,
  amount
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetPersonal :one
SELECT * FROM personal
WHERE id = $1 LIMIT 1;

-- name: Listpersonal :many
SELECT * FROM personal
ORDER BY id
LIMIT $1
OFFSET $2;
