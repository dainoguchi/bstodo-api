package handler

import (
	"github.com/dainoguchi/bstodo-api/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ハンドラーは実態に依存させていい気がする
type UserHandler struct {
	// usecaseのinterface。コンストラクタで実態を渡す
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	user, _ := h.usecase.GetByID(ctx, "1")

	type response struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	return c.JSON(http.StatusOK, response{
		ID:   user.ID,
		Name: user.Name,
	})
}
