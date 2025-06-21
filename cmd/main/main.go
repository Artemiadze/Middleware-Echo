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

	s.Use(MiddleWare) // Подключаем middleware
	// Задаём параметры обработчика
	// 127.0.0.1:8080/status
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

func MiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		val := c.Request().Header.Get("User-Role")
		if val == "admin" {
			log.Println("red button user detected")
		}

		err := next(c) // Вызываем следующий обработчик
		if err != nil {
			return err // Если произошла ошибка, возвращаем её
		}
		return nil // Возвращаем nil, если всё прошло успешно
	}
}
