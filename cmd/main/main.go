package main

import (
	"log"

	"github.com/Artemiadze/Middleware-Echo/internal/pkg/app"
)

func main() {
	app, err := app.New() // Создаём новое приложение
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run() // Запускаем приложение
	if err != nil {
		log.Fatal(err)
	}

}
