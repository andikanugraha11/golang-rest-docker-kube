package handlers

import (
	"github.com/andikanugraha11/golang-rest-docker-kube/models"
	"github.com/andikanugraha11/golang-rest-docker-kube/utils"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

func GetsAllUsers(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	var users []models.Users
	db.Find(&users)
	res := utils.JsonResponse(true,"success", users)
	return c.JSON(http.StatusOK, res)
}

func AddUser(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	var user models.Users
	if err := c.Bind(&user); err != nil {
		res := utils.JsonResponse(false,"error", err)
		return c.JSON(http.StatusOK, res)
	}
	db.Create(&user)
	res := utils.JsonResponse(true,"success", user.Id)
	return c.JSON(http.StatusOK, res)
}