package repository

import (
	"database/sql"
	"embed"
	"fmt"
	"main/util"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/jrivets/log4g"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type IDatabase interface {
	ConnectDatabase() *sql.DB
	RunMigration(embedMigrations embed.FS, Db *sql.DB)
}
type Database struct{}

func DatabaseInstance() *Database {
	return &Database{}
}

func (instance *Database) ConnectDatabase() *sql.DB {

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

	db, errSql := sql.Open("postgres", psqlSetup)
	errConn := db.Ping()
	if errSql != nil || errConn != nil {
		fmt.Println("There is an error while connecting to the database ", err, errSql, errConn)
		panic(err)
	}

	fmt.Println("Successfully connected to database!")

	return db
}

func (instance *Database) RunMigration(embedMigrations embed.FS, Db *sql.DB) {
	logger := log4g.GetLogger(util.LoggerName)

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		logger.Error(fmt.Sprintf("Error on set dialect %s", err.Error()))
		panic(err)
	}

	if err := goose.Up(Db, "migrations"); err != nil {
		logger.Error(fmt.Sprintf("Error on getting up the migrations %s", err.Error()))
		panic(err)
	}

}
