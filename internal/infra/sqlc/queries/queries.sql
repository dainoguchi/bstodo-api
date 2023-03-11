-- name: CreateTodo :one
INSERT INTO todos (
    id, title, description, done, priority, due_date, auth0_id
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         )
    RETURNING *;


-- name: ListTodos :many
SELECT * FROM todos
WHERE auth0_id = $1 OFFSET $2 LIMIT $3 ;

-- name: TotalTodoCount :one
SELECT count(id) FROM todos WHERE auth0_id = $1;

-- name: FindTodoByID :one
SELECT * FROM todos t WHERE t.id = $1;

-- name: FindTodoByAuth0ID :many
SELECT * FROM todos t WHERE t.auth0_id = $1;

-- name: DeleteTodo :exec
DELETE FROM todos t WHERE t.id = $1;

-- name: UpdateTodo :one
UPDATE todos
SET
    title = coalesce(sqlc.narg('title'), title),
    description = coalesce(sqlc.narg('description'), description),
    priority = coalesce(sqlc.narg('priority'), priority),
    due_date = coalesce(sqlc.narg('due_date'), due_date),
    done = coalesce(sqlc.narg('done'), done)
WHERE id = sqlc.arg('id')
    RETURNING *;
