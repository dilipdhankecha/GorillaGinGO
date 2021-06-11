package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	DriverName  string
	User        string
	Password    string
	Host        string
	Port        string
	DatabseName string
}

func NewDatabaseConnection() *DatabaseConfig {
	return &DatabaseConfig{
		DriverName:  "mysql",
		User:        "root",
		Password:    "smart",
		Host:        "localhost",
		Port:        "3306",
		DatabseName: "digpoc",
	}
}

func NewRepositoryConnection(config *DatabaseConfig) *sql.DB {
	connectionDetails := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DatabseName
	db, err := sql.Open(config.DriverName, connectionDetails)
	if err != nil {
		log.Fatal("Erorr Occur while connect with database..")
		log.Panic(err)
	}
	return db
}
