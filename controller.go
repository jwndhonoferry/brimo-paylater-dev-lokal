package main

import (
	"log"
	"os"
	"strings"
)

func GetBrimoDataPhoenixDB(req BrimoPaylaterDataByAcctNoRequest) ([]BrimoPaylaterDataResponse, error) {
	var BrimoDataResponseBody []BrimoPaylaterDataResponse
	newAcctno := strings.TrimLeft(req.AcctNo, "0")
	// fmt.Println("MAIN DB NYAAAA", os.Getenv("MAIN_DB"))
	if os.Getenv("MAIN_DB") == "hbase" {
		db, _ := Connect()
		err = db.
			Table("CUSTOMER_CERIA_CS").
			Select("*").
			Where("ACCOUNT_NUMBER = ?", newAcctno).
			Scan(&BrimoDataResponseBody).Error
	} else {
		db := getPostgresDb()
		// err = db.Raw("SELECT * FROM brimo_pl_dev.brimo_paylater_whitelist_dev WHERE acctno = ? ", newAcctno).Scan(&BrimoDataResponseBody).Error
		err = db.
<<<<<<< HEAD
			Table(`brimo_pl_dev.brimo_paylater_whitelist_dev`).
			Select("*").
			Where(`acctno = ?`, newAcctno).
			Scan(&BrimoDataResponseBody).Error
=======
			Table(`"public.brimo_paylater_whitelist"`).
			Select("*").
			Where(`"acctno" = ?`, newAcctno).
			Scan(&ceriaDataResponseBody).Error
>>>>>>> 26312a2e32b0152e2754a471b4e1d8b8a9fe0186
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return BrimoDataResponseBody, nil
}

// func GetCeriaFeatureData(req BrimoPaylaterDataByAcctNoRequest) ([]BrimoPaylaterDataByAcctNoRequest, error) {
// 	var ceriaDataResponseBody []BrimoPaylaterDataByAcctNoRequest
// 	newAcctno := strings.TrimLeft(req.AcctNo, "0")

// 	if os.Getenv("MAIN_DB") == "hbase" {
// 		db, _ := Connect()
// 		err = db.
// 			Table("CS_CERIA_IDX_FEATURE").
// 			Select("*").
// 			Where("ACCTNO = ?", newAcctno).
// 			Scan(&ceriaDataResponseBody).Error
// 	} else {
// 		db := getPostgresDb()
// 		err = db.
// 			Table(`"CERIA.CS_CERIA_IDX_FEATURE"`).
// 			Select("*").
// 			Where(`"ACCTNO" = ?`, newAcctno).
// 			Scan(&ceriaDataResponseBody).Error
// 	}

// 	return ceriaDataResponseBody, nil
// }

// func GetCeriaDataByEmailPhoenixDBV2(req BrimoPaylaterDataByAcctNoRequest) ([]BrimoPaylaterDataByAcctNoRequest, error) {
// 	var ceriaDataByEmailResponseBody []BrimoPaylaterDataByAcctNoRequest

// 	//Connect To HBASE Phoenix DB
// 	var db, err = Connect()
// 	//Check Connection
// 	if err != nil {
// 		fmt.Println("1")
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}

// 	//Get Data From Phoenix Table
// 	err = db.
// 		Table("CUSTOMER_CERIA_CS_BY_MAIL_V2").
// 		Select("*").
// 		Where("EMAIL = ?", strings.ToLower(req.Email)).
// 		Scan(&ceriaDataByEmailResponseBody).Error
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}

// 	return ceriaDataByEmailResponseBody, nil
// }

// func GetCeriaDataByPhonePhoenixDBV2(req CeriaDataByPhoneRequest) ([]CeriaDataByPhoneResponseV2, error) {
// 	var ceriaDataByPhoneResponseBody []CeriaDataByPhoneResponseV2

// 	//Connect To HBASE Phoenix DB
// 	var db, err = Connect()
// 	//Check Connection
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}

// 	//Get Data From Phoenix Table
// 	err = db.
// 		Table("CUSTOMER_CERIA_CS_BY_PHONE_V2").
// 		Select("*").
// 		Where("HANDPHONE = ?", strings.ToLower(req.Handphone)).
// 		Scan(&ceriaDataByPhoneResponseBody).Error
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}

// 	return ceriaDataByPhoneResponseBody, nil
// }
