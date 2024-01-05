package main

import (
	"github.com/Ericrulec/gohtmx/handler"
	"github.com/labstack/echo"
)

func main() {
	app := echo.New()

	indexHandler := handler.IndexHandler{}
	app.Static("/", "assets")

	app.GET("/", indexHandler.GetIndex)

	err := app.Start(":3000")
	if err != nil {
		panic(err)
	}
}
