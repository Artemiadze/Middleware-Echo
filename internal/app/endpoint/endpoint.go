package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

// Endpoint - структура, в которой хранится сервис
type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

// Handler - функция, в которой заключена логика ответа на запрос
func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DaysLeft() // Вызываем метод DaysLeft из сервиса

	s := fmt.Sprintf("Days left 1.01.2027: %d\n", d)

	err := ctx.String(http.StatusOK, s) // Возвращаем статус 200 OK
	if err != nil {
		// Если произошла ошибка при отправке ответа, логируем её
		return err
	}
	return nil
}
