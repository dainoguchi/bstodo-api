package usecase

import (
	"context"
	"database/sql"
	"github.com/dainoguchi/bstodo-api/internal/entity"
)

type UserUsecase interface {
	GetByID(context.Context, string) (*entity.User, error)
}

func NewUserUsecase(db *sql.DB) UserUsecase {
	return &userUsecase{db: db}
}

type userUsecase struct {
	db *sql.DB
}

// 実装はまだ
func (u userUsecase) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return &entity.User{ID: id, Name: "テスト君"}, nil
}
