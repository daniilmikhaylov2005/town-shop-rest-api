package main

import (
	"github.com/daniilmikhaylov2005/town-shop-rest-api/handlers"
	m "github.com/daniilmikhaylov2005/town-shop-rest-api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	auth := e.Group("/auth")
	auth.POST("/signup", handlers.Signup)
	auth.POST("/signin", handlers.Signin)

	api := e.Group("/api")
	api.GET("/goods/:category", handlers.GetAllGoods)
	api.GET("/goods/:category/:id", handlers.GetGoodById)

	admin := e.Group("/admin")
	admin.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		ParseTokenFunc: m.ParseToken,
	}))
	admin.POST("/good", handlers.InsertGood)
	admin.PUT("/good/:id", handlers.UpdateGood)
	admin.DELETE("/good/:id", handlers.DeleteGood)

	e.Logger.Fatal(e.Start(":8000"))
}
