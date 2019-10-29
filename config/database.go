package config

import (
	"os"
	"strconv"
)

type Database struct {
	DbHost     	string
	DbPort     	int
	DbUser     	string
	DbPassword 	string
	DbName     	string
	DbMaxCon   	int
	DbTimeOut	int
}

func DatabaseEnv() (Database) {
	db := Database{}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		dbHost = "localhost"
	}
	db.DbHost = dbHost

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		dbPort = "5432"
	}
	db.DbPort, _ = strconv.Atoi(dbPort)

	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		dbUser = "dummy"
	}
	db.DbUser = dbUser

	dbPass, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		dbPass = "dummy"
	}
	db.DbPassword = dbPass

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		dbName = "dummy"
	}
	db.DbName = dbName

	dbMaxcon, ok := os.LookupEnv("DB_MAX_CON")
	if !ok {
		dbMaxcon = "-1"
	}
	y, _ := strconv.Atoi(dbMaxcon)
	db.DbMaxCon = y

	dbTimeout, ok := os.LookupEnv("DB_TIME_OUT")
	if !ok {
		dbTimeout = "0"
	}
	dbTimeoutint, _ := strconv.Atoi(dbTimeout)
	db.DbTimeOut = dbTimeoutint


	return db
}