-- name: CreateTodo :one
INSERT INTO todos (
    id, title, description, done, priority, due_date, auth0_id
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         )
    RETURNING *;
