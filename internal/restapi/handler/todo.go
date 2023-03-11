package handler

import (
	"github.com/dainoguchi/bstodo-api/internal/restapi/httputil"
	"github.com/dainoguchi/bstodo-api/internal/usecase"
	"github.com/dainoguchi/bstodo-api/internal/usecase/input"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type TodoHandler struct {
	// usecaseのinterface。コンストラクタで実態を渡す
	uc usecase.TodoUsecase
}

func NewTodoHandler(uc usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{uc: uc}
}

func (h *TodoHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	todo, err := h.uc.GetTodo(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// response

	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) List(c echo.Context) error {
	var req struct {
		Offset int `json:"offset" query:"offset"`
		Limit  int `json:"limit" query:"limit"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	auth0ID := httputil.GetAuth0ID(ctx)

	todoList, err := h.uc.ListTodos(ctx, auth0ID, req.Limit, req.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// responseに詰め替え

	return c.JSON(http.StatusOK, todoList)
}

func (h *TodoHandler) Create(c echo.Context) error {
	var req struct {
		Title       string     `json:"title"`
		Priority    string     `json:"priority"`
		Description *string    `json:"description"`
		DueDate     *time.Time `json:"due_date"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	auth0ID := httputil.GetAuth0ID(ctx)

	todo, err := h.uc.CreateTodo(ctx, &input.CreateTodoInput{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		Auth0ID:     auth0ID,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// response 詰め替え
	// sqlcのresponseにomitemptyをつけるオプションが存在しなかったのです,,,
	res := struct {
		ID          uuid.UUID  `json:"id"`
		Title       string     `json:"title"`
		Description *string    `json:"description,omitempty"`
		Done        bool       `json:"done"`
		Priority    string     `json:"priority"`
		DueDate     *time.Time `json:"due_date,omitempty"`
		Auth0ID     string     `json:"auth0_id"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
	}{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Done:        todo.Done,
		Priority:    todo.Priority,
		DueDate:     todo.DueDate,
		Auth0ID:     todo.Auth0ID,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *TodoHandler) Update(c echo.Context) error {

	var req struct {
		ID          string     `json:"id" param:"id"`
		Title       *string    `json:"title"`
		Priority    *string    `json:"priority"`
		Description *string    `json:"description"`
		DueDate     *time.Time `json:"due_date"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	auth0ID := httputil.GetAuth0ID(ctx)

	todo, err := h.uc.UpdateTodo(ctx, &input.UpdateTodoInput{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		Auth0ID:     auth0ID,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// response 詰め替え
	// sqlcのresponseにomitemptyをつけるオプションが存在しなかったのです,,,
	res := struct {
		ID          uuid.UUID  `json:"id"`
		Title       string     `json:"title"`
		Description *string    `json:"description,omitempty"`
		Done        bool       `json:"done"`
		Priority    string     `json:"priority"`
		DueDate     *time.Time `json:"due_date,omitempty"`
		Auth0ID     string     `json:"auth0_id"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
	}{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Done:        todo.Done,
		Priority:    todo.Priority,
		DueDate:     todo.DueDate,
		Auth0ID:     todo.Auth0ID,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *TodoHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	deletedID, err := h.uc.DeleteTodo(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// response
	res := struct {
		ID uuid.UUID `json:"id"`
	}{
		ID: *deletedID,
	}

	return c.JSON(http.StatusOK, res)
}
