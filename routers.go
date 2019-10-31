package main

import (
	"github.com/andikanugraha11/golang-rest-docker-kube/handlers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

func dataSourceMiddleware(dataStore *gorm.DB)  echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", dataStore)
			return next(c)
		}
	}
}

func UserRouter(e *echo.Echo, db *gorm.DB)  {
	// Set Group
	userRoute := e.Group("/user")
	//userRoute.GET("", homePage)
	userRoute.GET("", handlers.GetsAllUsers, dataSourceMiddleware(db))
	userRoute.POST("", handlers.AddUser)
}

func homePage(c echo.Context) error {
	res := "Helloworld"
	return c.JSON(http.StatusOK, res)
}
