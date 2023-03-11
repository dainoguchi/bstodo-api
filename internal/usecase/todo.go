package usecase

import (
	"context"
	"fmt"
	"github.com/dainoguchi/bstodo-api/internal/infra/postgres"
	"github.com/dainoguchi/bstodo-api/internal/infra/sqlc"
	"github.com/dainoguchi/bstodo-api/internal/usecase/input"
	"github.com/google/uuid"
)

type TodoUsecase interface {
	CreateTodo(ctx context.Context, input *input.CreateTodoInput) (*sqlc.Todo, error)
	GetTodo(ctx context.Context, id string) (*sqlc.Todo, error)
	ListTodos(ctx context.Context, auth0ID string, limit int, offset int) (*ListTodoResponse, error)
	UpdateTodo(ctx context.Context, input *input.UpdateTodoInput) (*sqlc.Todo, error)
	DeleteTodo(ctx context.Context, id string) (*uuid.UUID, error)
}

// テスト時transaction持たせてrollbackしたい為, pgx.Connとpgx.Txを抽象化したWrapper構造体を定義
func NewTodoUsecase(db postgres.PgxWrapper) TodoUsecase {
	return &todoUsecase{db: db}
}

type todoUsecase struct {
	db postgres.PgxWrapper
}

func (t *todoUsecase) CreateTodo(ctx context.Context, input *input.CreateTodoInput) (*sqlc.Todo, error) {
	q := sqlc.New(t.db)

	if err := input.Validate(); err != nil {
		return nil, err
	}

	todo, err := q.CreateTodo(ctx, sqlc.CreateTodoParams{
		ID:          uuid.New(),
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
		Done:        false,
		Priority:    input.Priority,
		Auth0ID:     input.Auth0ID,
	})

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *todoUsecase) GetTodo(ctx context.Context, id string) (*sqlc.Todo, error) {
	q := sqlc.New(t.db)

	// ここでuuid使って良いのかなー
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// todoでいいのかな
	todo, err := q.FindTodoByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

type ListTodoResponse struct {
	TotalCount int
	Todos      []*sqlc.Todo
}

// この関心毎もまとめた方が良さげではあるよね
// auth0の型にそってるかもチェックするコード仕込んだ方が良さげ
const defaultLimit int = 10

// optionでまとめた方が良さそーww
func (t *todoUsecase) ListTodos(ctx context.Context, auth0ID string, limit int, offset int) (*ListTodoResponse, error) {
	q := sqlc.New(t.db)

	// もし空なら最低でも一つは
	if limit == 0 {
		limit = defaultLimit
	}

	// todoでいいのかな
	todos, err := q.ListTodos(ctx, sqlc.ListTodosParams{
		Auth0ID: auth0ID,
		Limit:   int32(limit),
		Offset:  int32(offset),
	})

	if err != nil {
		return nil, fmt.Errorf("ListTodos: %w", err)
	}

	totalCount, err := q.TotalTodoCount(ctx, auth0ID)
	if err != nil {
		return nil, fmt.Errorf("ListTodos: %w", err)
	}

	return &ListTodoResponse{
		TotalCount: int(totalCount),
		Todos:      todos,
	}, nil
}

func (t *todoUsecase) UpdateTodo(ctx context.Context, input *input.UpdateTodoInput) (*sqlc.Todo, error) {
	q := sqlc.New(t.db)

	// 欲しいデータが分からない
	if err := input.Validate(); err != nil {
		return nil, err
	}

	uid, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, err
	}

	// 怪しい
	todo, err := q.UpdateTodo(ctx, sqlc.UpdateTodoParams{
		ID:          uid,
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
		Done:        input.Done,
		Priority:    input.Priority,
	})

	fmt.Println(todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *todoUsecase) DeleteTodo(ctx context.Context, id string) (*uuid.UUID, error) {
	q := sqlc.New(t.db)

	// ここでuuid使って良いのかなー
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// todoでいいのかな
	err = q.DeleteTodo(ctx, uid)
	if err != nil {
		return nil, err
	}

	// 検証のためにid返したいけど...
	return &uid, nil
}
