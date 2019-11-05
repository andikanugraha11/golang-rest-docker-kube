package main

import (
	"github.com/andikanugraha11/golang-rest-docker-kube/handlers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
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

	// Router Handler
	userRoute.GET("", handlers.GetsAllUsers, dataSourceMiddleware(db))
	userRoute.POST("", handlers.AddUser, dataSourceMiddleware(db))
}
