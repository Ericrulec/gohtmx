package main

import (
	"github.com/Ericrulec/gohtmx/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/", "assets")

	indexHandler := handler.IndexHandler{}
	postsHandler := handler.PostsHandler{}
	contentHandler := handler.ContentHandler{}

	app.GET("/", indexHandler.GetIndex)
	app.GET("/posts", postsHandler.GetPosts)
	app.GET("/posts/passwords", contentHandler.GetPasswordsPage)

	err := app.Start(":3000")
	if err != nil {
		panic(err)
	}
}
