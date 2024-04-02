package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDatabase() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}

	//Load values from .env file
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")

	// set up postgres sql to open it.
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)

	fmt.Println(psqlSetup)

	db, errSql := sql.Open("postgres", psqlSetup)
	errConn := db.Ping()
	if errSql != nil || errConn != nil {
		fmt.Println("There is an error while connecting to the database ", err, errSql, errConn)
		panic(err)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}
}
