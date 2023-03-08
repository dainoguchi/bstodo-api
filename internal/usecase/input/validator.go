package input

import (
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()

	// 参考 https://qiita.com/RunEagler/items/ad79fc860c3689797ccc
	v.RegisterValidation("is_priority", isPriority)
}

const (
	PriorityHigh = "high"
	PriorityMid  = "mid"
	PriorityLow  = "low"
)

type Priority string

func (p Priority) Validate() bool {
	switch p {
	case PriorityHigh, PriorityMid, PriorityLow:
		return true
	default:
		return false
	}
}

func isPriority(fl validator.FieldLevel) bool { //引数の型、返り値は固定
	p := fl.Field().String()

	// nullの場合はdefault valueとしてmidが入る為許可する　ができない、、、
	if Priority(p).Validate() {
		return true
	}
	return false
}
