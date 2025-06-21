package app

import (
	"fmt"
	"log"

	"github.com/Artemiadze/Middleware-Echo/internal/app/endpoint"
	"github.com/Artemiadze/Middleware-Echo/internal/app/middleware"
	"github.com/Artemiadze/Middleware-Echo/internal/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}             // Создаём новый экземпляр приложения
	a.s = service.New()     // Создаём новый сервис
	a.e = endpoint.New(a.s) // Создаём новый эндпоинт с сервисом

	a.echo = echo.New()

	a.echo.Use(middleware.MiddleWare) // Подключаем middleware
	// Задаём параметры обработчика
	// 127.0.0.1:8080/status
	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server is running...")

	err := a.echo.Start(":8080")
	if err != nil {
		// команда s.Start может возвращать ошибку
		log.Fatal(err)
	}

	return nil
}
