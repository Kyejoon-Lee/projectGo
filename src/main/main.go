package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"api/handlers"
	"controller"

)


func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080","http://localhost:8081","http://localhost:63342","http://52.78.172.184:8080","http://52.78.172.184:8081"},
		AllowHeaders: []string{echo.HeaderOrigin,echo.HeaderContentType,echo.HeaderAccept,"Authorization"},
	}))
	e.GET("/",controller.MainPage)
	e.GET("/profits/:id",handlers.GetData)
	e.Start(":8080")

}
