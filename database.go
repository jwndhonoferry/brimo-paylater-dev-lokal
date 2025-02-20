package main

import (
	"log"

	_ "github.com/apache/calcite-avatica-go/v5"
	"github.com/jinzhu/gorm"
	gormio "gorm.io/gorm"
)

// Db variable
var Db *gorm.DB
var PGDb *gormio.DB
var err error

func InitDb() {
	// Db, err = connectHBase()
	// if err != nil {
	// 	panic(err)
	// }

	PGDb, err = connectPostgres()
	if err != nil {
		panic(err)
	}
}

func getPostgresDb() *gormio.DB {
	if PGDb == nil {
		log.Fatal("PostgreSQL connection is not initialized")
	}
	return PGDb
}
