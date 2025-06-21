package middleware

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func MiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		val := c.Request().Header.Get("User-Role")
		// проверяем, есть ли в заголовке User-Role значение "admin", без учёта регистра
		if strings.EqualFold(val, roleAdmin) {
			log.Println("red button user detected")
		}

		err := next(c) // Вызываем следующий обработчик
		if err != nil {
			return err // Если произошла ошибка, возвращаем её
		}
		return nil // Возвращаем nil, если всё прошло успешно
	}
}
