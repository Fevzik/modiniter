package modiniter

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const (
	codeKey = "code"
	dataKey = "data"
	failKey = "fail"

	performedKey = "performed"
)

type DataFunc func() (interface{}, error)
type ErrFunc func() error

type Result interface {
	SetCode(code int) Result
	SetData(data interface{}) Result
	SetFail(fail error) Result
	Performed() Result
	Map() fiber.Map
}

type result struct {
	data fiber.Map
}

func New() Result {
	res := result{data: fiber.Map{
		codeKey: http.StatusOK,
	}}
	return res
}

func (r result) SetCode(code int) Result {
	r.data[codeKey] = code
	return r
}

func (r result) SetData(data interface{}) Result {
	r.data[dataKey] = data
	return r
}

func (r result) Performed() Result {
	r.data[dataKey] = fiber.Map{performedKey: true}
	return r
}

func (r result) SetFail(fail error) Result {
	r.data[failKey] = fail
	return r
}

func (r result) Map() fiber.Map {
	return r.data
}
