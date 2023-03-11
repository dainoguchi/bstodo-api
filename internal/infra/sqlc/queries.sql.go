// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: queries.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (
    id, title, description, done, priority, due_date, auth0_id
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         )
    RETURNING id, title, description, done, priority, due_date, auth0_id, created_at, updated_at
`

type CreateTodoParams struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Done        bool       `json:"done"`
	Priority    string     `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	Auth0ID     string     `json:"auth0_id"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (*Todo, error) {
	row := q.db.QueryRow(ctx, createTodo,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Done,
		arg.Priority,
		arg.DueDate,
		arg.Auth0ID,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Done,
		&i.Priority,
		&i.DueDate,
		&i.Auth0ID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos t WHERE t.id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteTodo, id)
	return err
}

const findTodoByAuth0ID = `-- name: FindTodoByAuth0ID :many
SELECT id, title, description, done, priority, due_date, auth0_id, created_at, updated_at FROM todos t WHERE t.auth0_id = $1
`

func (q *Queries) FindTodoByAuth0ID(ctx context.Context, auth0ID string) ([]*Todo, error) {
	rows, err := q.db.Query(ctx, findTodoByAuth0ID, auth0ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Done,
			&i.Priority,
			&i.DueDate,
			&i.Auth0ID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findTodoByID = `-- name: FindTodoByID :one
SELECT id, title, description, done, priority, due_date, auth0_id, created_at, updated_at FROM todos t WHERE t.id = $1
`

func (q *Queries) FindTodoByID(ctx context.Context, id uuid.UUID) (*Todo, error) {
	row := q.db.QueryRow(ctx, findTodoByID, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Done,
		&i.Priority,
		&i.DueDate,
		&i.Auth0ID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, description, done, priority, due_date, auth0_id, created_at, updated_at FROM todos
WHERE auth0_id = $1 OFFSET $2 LIMIT $3
`

type ListTodosParams struct {
	Auth0ID string `json:"auth0_id"`
	Offset  int32  `json:"offset"`
	Limit   int32  `json:"limit"`
}

func (q *Queries) ListTodos(ctx context.Context, arg ListTodosParams) ([]*Todo, error) {
	rows, err := q.db.Query(ctx, listTodos, arg.Auth0ID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Done,
			&i.Priority,
			&i.DueDate,
			&i.Auth0ID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const totalTodoCount = `-- name: TotalTodoCount :one
SELECT count(id) FROM todos WHERE auth0_id = $1
`

func (q *Queries) TotalTodoCount(ctx context.Context, auth0ID string) (int64, error) {
	row := q.db.QueryRow(ctx, totalTodoCount, auth0ID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET
    title = coalesce($1, title),
    description = coalesce($2, description),
    priority = coalesce($3, priority),
    due_date = coalesce($4, due_date),
    done = coalesce($5, done)
WHERE id = $6
    RETURNING id, title, description, done, priority, due_date, auth0_id, created_at, updated_at
`

type UpdateTodoParams struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Priority    *string    `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	Done        *bool      `json:"done"`
	ID          uuid.UUID  `json:"id"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (*Todo, error) {
	row := q.db.QueryRow(ctx, updateTodo,
		arg.Title,
		arg.Description,
		arg.Priority,
		arg.DueDate,
		arg.Done,
		arg.ID,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Done,
		&i.Priority,
		&i.DueDate,
		&i.Auth0ID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
