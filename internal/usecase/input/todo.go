package input

import "time"

type CreateTodoInput struct {
	Auth0ID string `validate:"required"`
	// 最低１文字以上
	Title string `validate:"required,gte=1"'`
	// 参照外部キー制約を表す
	Priority string `validate:"is_priority"'`

	Description *string
	DueDate     *time.Time
}

func (in *CreateTodoInput) Validate() error {
	if err := v.Struct(in); err != nil {
		return err
	}

	return nil
}

type UpdateTodoInput struct {
	ID      string `validate:"required,uuid4"`
	Auth0ID string `validate:"required"`
	// あるなら1文字以上。nilも可
	Title       *string `validate:"omitempty,gte=1"'`
	Priority    *string `validate:"is_priority"'`
	Done        *bool
	Description *string
	DueDate     *time.Time
}

func (in *UpdateTodoInput) Validate() error {
	if err := v.Struct(in); err != nil {
		return err
	}

	return nil
}
