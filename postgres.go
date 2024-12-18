package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const database = "PostgreSQL Server"

func connectPostgres() (*gorm.DB, error) {
	var (
		err       error
		dialector = os.Getenv("POSTGRES_DIALECTOR")
	)

	PGDb, err = gorm.Open(postgres.Open(dialector), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Println("[POSTGRES] ", err.Error())
		return nil, err
	}

	return PGDb.Debug(), err
}
