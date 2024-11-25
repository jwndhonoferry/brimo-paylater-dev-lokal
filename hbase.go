package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/apache/calcite-avatica-go/v5"
	"github.com/jinzhu/gorm"
)

func connectHBase() (*gorm.DB, error) {
	var conn string

	if os.Getenv("AUTH_MODE") == "kerberos" {
		pqs_host := os.Getenv("PQS_HOST")
		balance_principal := os.Getenv("PRINCIPAL")
		balance_keytab := os.Getenv("KEYTAB")

		conn = fmt.Sprintf("%s/%s?authentication=%s&principal=%s&keytab=%s&krb5Conf=%s&krb5CredentialsCache=%s", pqs_host, os.Getenv("HBASE_SCHEMA"), os.Getenv("AUTH"), balance_principal, balance_keytab, os.Getenv("KRB5"), os.Getenv("KRB5_CACHE"))
	} else {
		conn = os.Getenv("PHOENIX_DB_URL")
	}
	Db, err = gorm.Open("avatica", conn)

	if err != nil {
		log.Println("Connection Database Error ", err.Error())
	} else {
		Db.DB().SetMaxIdleConns(5)
		Db.DB().SetMaxOpenConns(20)
		Db.LogMode(true)
		log.Println("Connected")
	}

	return Db, err
}

func Connect() (*gorm.DB, error) {
	return Db, err
}
