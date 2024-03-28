package config

import (
	"embed"
	"fmt"
	repository "main/repositories"
	"main/util"

	"github.com/jrivets/log4g"
	"github.com/pressly/goose/v3"
)

func RunMigration(embedMigrations embed.FS) {
	logger := log4g.GetLogger(util.LoggerName)
	repository.ConnectDatabase()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		logger.Error(fmt.Sprintf("Error on set dialect %s", err.Error()))
		panic(err)
	}

	if err := goose.Up(repository.Db, "migrations"); err != nil {
		logger.Error(fmt.Sprintf("Error on getting up the migrations %s", err.Error()))
		panic(err)
	}

}
