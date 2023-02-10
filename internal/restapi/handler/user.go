package handler

import (
	"encoding/json"
	"github.com/dainoguchi/bstodo-api/internal/usecase"
	"log"
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

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, _ := h.usecase.GetByID(ctx, "1")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
