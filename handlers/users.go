package handlers

import (
	"github.com/andikanugraha11/golang-rest-docker-kube/models"
	"github.com/andikanugraha11/golang-rest-docker-kube/utils"
	"github.com/labstack/echo"
	"net/http"
)

func UserDetail(c echo.Context) error {
	//user := map[string]string {"hello" : "Hello"}
	userDet := models.Users{}
	userDet.Name = "Andika Nugraha"

	//js, err := json.Marshal(profile)

	res := utils.JsonResponse(true,"success", userDet)
	return c.JSON(http.StatusOK, res)
}