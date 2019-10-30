package main

import (
	"fmt"
	"github.com/andikanugraha11/golang-rest-docker-kube/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"os"
)

var conf config.AppConfig

func handleRequest(){
	log.Println("Starting server at 8080")
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${status}]-${method}-${uri}\n",
	}))

	// Register Router
	UserRouter(e)

	e.Start(":8080")
}

func main() {
	// Read Env Variable
	conf, err := config.ReadEnvVar()
	if err != nil {
		log.Panicf("Fail read environtment variable : %v", err)
	}

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

	// Get all parameter on running app
	if len(os.Args) > 1 {
		commandParams := os.Args[1]
		if commandParams == "migrate" {
			MigrateDB(db)
		}
	}

	handleRequest()
}
