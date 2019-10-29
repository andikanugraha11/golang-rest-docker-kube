package main

import (
	"github.com/andikanugraha11/golang-rest-docker-kube/config"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/andikanugraha11/golang-rest-docker-kube/models"
)

func main()  {
	conf, err := config.ReadEnvVar()
	// DB init with gorm
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Database.DbHost, conf.Database.DbPort,
		conf.Database.DbUser, conf.Database.DbPassword, conf.Database.DbName,
	)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Panicf("Fail connect to database : %v", err)
	}
	defer db.Close()

	log.Println("Start migrate ...")
	db.AutoMigrate(&models.Users{})
	log.Println("Finish migrate ...")
}