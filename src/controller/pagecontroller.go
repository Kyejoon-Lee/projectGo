package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func MainPage(e echo.Context) error{
	return e.Redirect(http.StatusTemporaryRedirect,"http://52.78.172.184:8081/index.html")
}
