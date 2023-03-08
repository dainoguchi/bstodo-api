package usecase

import (
	"context"
	"github.com/dainoguchi/bstodo-api/internal/infra/postgres"
	"github.com/dainoguchi/bstodo-api/internal/infra/sqlc"
	"github.com/dainoguchi/bstodo-api/internal/usecase/input"
	"github.com/google/uuid"
)

type TodoUsecase interface {
	CreateTodo(ctx context.Context, input *input.CreateTodoInput) (*sqlc.Todo, error)
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
