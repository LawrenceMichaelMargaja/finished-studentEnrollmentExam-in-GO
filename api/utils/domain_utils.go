package utils

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// initializeClient establishes a database connection for MYSQL
func GetMYSQLConnection() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	connString := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":"+ dbPort +")/" + dbDatabase + "?parseTime=true"
	client, err := sqlx.Open("mysql", connString)
	if err != nil {
		log.Println(err)
		panic("Something went wrong with your database connection.")
	}

	maxConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	client.SetMaxOpenConns(maxConnections)

	return client
}