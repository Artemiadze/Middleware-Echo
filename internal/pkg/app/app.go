package app

import (
	"github.com/Artemiadze/Middleware-Echo/internal/app/endpoint"
	"github.com/Artemiadze/Middleware-Echo/internal/app/service"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
}

func New() *App {
	a := &App{}              // Создаём новый экземпляр приложения
	a.s := service.New()     // Создаём новый сервис
	a.e := endpoint.New(a.s) // Создаём новый эндпоинт с сервисом
}
