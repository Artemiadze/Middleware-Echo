package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Server is running...")

	s := echo.New()
	// Задаём параметры обработчика
	s.GET("/status", Handler)

	err := s.Start(":8080")
	if err != nil {
		// команда s.Start может возвращать ошибку
		log.Fatal(err)
		return
	}
}

// Handler - функция, в которой заключена логика ответа на запрос
func Handler(c echo.Context) error {
	// обьект дата
	d := time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d) // Вычисляем время до указанной даты
	s := fmt.Sprintf("Времени до 1.01.2027: %d\n", int64(dur.Hours()/24))

	err := c.String(http.StatusOK, s) // Возвращаем статус 200 OK
	if err != nil {
		// Если произошла ошибка при отправке ответа, логируем её
		return err
	}
	return nil
}
