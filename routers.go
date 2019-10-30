package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/andikanugraha11/golang-rest-docker-kube/handlers"
)

func UserRouter(e *echo.Echo)  {
	// Set Group
	userRoute := e.Group("/user")

	userRoute.GET("", homePage)
	userRoute.GET("/detail", handlers.UserDetail)
}

func homePage(c echo.Context) error {
	res := "Helloworld"
	return c.JSON(http.StatusOK, res)
}
