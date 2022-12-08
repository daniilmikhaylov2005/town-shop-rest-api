package main

import (
	"github.com/daniilmikhaylov2005/town-shop-rest-api/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	auth := e.Group("/auth")

	auth.POST("/signup", handlers.Signup)
	auth.POST("/signin", handlers.Signin)

	api := e.Group("/api")
	api.GET("/goods/:category", handlers.GetAllGoods)
	api.GET("/goods/:category/:id", handlers.GetGoodById)

	e.Logger.Fatal(e.Start(":8000"))
}
