package main

import (
	"github.com/Ericrulec/gohtmx/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()
	//e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("ejen.no")
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	// e.Pre(middleware.HTTPSWWWRedirect())

	e.Static("/", "assets")

	indexHandler := handler.IndexHandler{}
	postsHandler := handler.PostsHandler{}
	contentHandler := handler.ContentHandler{}

	e.GET("/", indexHandler.GetIndex)
	e.GET("/posts", postsHandler.GetPosts)

	e.GET("/posts/passwords", contentHandler.GetPasswordsPage)

	//e.Logger.Fatal(e.StartAutoTLS(":3000"))
	e.Logger.Fatal(e.Start(":3000"))
}
