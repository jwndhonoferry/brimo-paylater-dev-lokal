package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetCeriaDataPhoenixDB(req CeriaDataRequest) ([]CeriaDataResponse, error) {
	var ceriaDataResponseBody []CeriaDataResponse
	newAcctno := strings.TrimLeft(req.AcctNo, "0")

	if os.Getenv("MAIN_DB") == "hbase" {
		db, _ := Connect()
		err = db.
			Table("CUSTOMER_CERIA_CS").
			Select("*").
			Where("ACCOUNT_NUMBER = ?", newAcctno).
			Scan(&ceriaDataResponseBody).Error
	} else {
		db := getPostgresDb()
		err = db.
			Table(`"CERIA.CUSTOMER_CERIA_CS"`).
			Select("*").
			Where(`"ACCOUNT_NUMBER" = ?`, newAcctno).
			Scan(&ceriaDataResponseBody).Error
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return ceriaDataResponseBody, nil
}

func GetCeriaFeatureData(req CeriaDataRequest) ([]CeriaFeatureDataResponse, error) {
	var ceriaDataResponseBody []CeriaFeatureDataResponse
	newAcctno := strings.TrimLeft(req.AcctNo, "0")

	if os.Getenv("MAIN_DB") == "hbase" {
		db, _ := Connect()
		err = db.
			Table("CS_CERIA_IDX_FEATURE").
			Select("*").
			Where("ACCTNO = ?", newAcctno).
			Scan(&ceriaDataResponseBody).Error
	} else {
		db := getPostgresDb()
		err = db.
			Table(`"CERIA.CS_CERIA_IDX_FEATURE"`).
			Select("*").
			Where(`"ACCTNO" = ?`, newAcctno).
			Scan(&ceriaDataResponseBody).Error
	}

	return ceriaDataResponseBody, nil
}

func GetCeriaDataByEmailPhoenixDBV2(req CeriaDataByEmailRequest) ([]CeriaDataByEmailResponseV2, error) {
	var ceriaDataByEmailResponseBody []CeriaDataByEmailResponseV2

	//Connect To HBASE Phoenix DB
	var db, err = Connect()
	//Check Connection
	if err != nil {
		fmt.Println("1")
		fmt.Println(err.Error())
		return nil, err
	}

	//Get Data From Phoenix Table
	err = db.
		Table("CUSTOMER_CERIA_CS_BY_MAIL_V2").
		Select("*").
		Where("EMAIL = ?", strings.ToLower(req.Email)).
		Scan(&ceriaDataByEmailResponseBody).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return ceriaDataByEmailResponseBody, nil
}

func GetCeriaDataByPhonePhoenixDBV2(req CeriaDataByPhoneRequest) ([]CeriaDataByPhoneResponseV2, error) {
	var ceriaDataByPhoneResponseBody []CeriaDataByPhoneResponseV2

	//Connect To HBASE Phoenix DB
	var db, err = Connect()
	//Check Connection
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	//Get Data From Phoenix Table
	err = db.
		Table("CUSTOMER_CERIA_CS_BY_PHONE_V2").
		Select("*").
		Where("HANDPHONE = ?", strings.ToLower(req.Handphone)).
		Scan(&ceriaDataByPhoneResponseBody).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return ceriaDataByPhoneResponseBody, nil
}
