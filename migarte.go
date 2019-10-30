package main

import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/andikanugraha11/golang-rest-docker-kube/models"
)

func MigrateDB(db *gorm.DB)  {
	log.Println("Start migrate ...")
	db.AutoMigrate(&models.Users{})
	log.Println("Finish migrate ...")
}