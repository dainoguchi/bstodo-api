-- name: ListUsers :many
SELECT * FROM users;

-- name: FindUser :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
    id, auth0_id, email, name
) VALUES (
             $1, $2, $3, $4
         )
    RETURNING id;

