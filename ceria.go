package main

import (
	"net/http"
	"whitelist_ceria/utils"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

// Ceria Data Hadler, Get Ceria Data, Using Method GET and URL Param
func BrimoData(c *gin.Context) {
	//Call Models Data Request
	var BrimomDataRequestBody BrimoPaylaterDataByAcctNoRequest

	tx := apm.DefaultTracer.StartTransaction("Get Brimo Paylater Whitelist by Acct No", "request")
	defer tx.End()

	//Read Acctno Request
	accnum := c.Param("accnum")

	//Append Acctno Request To Models
<<<<<<< HEAD
	BrimomDataRequestBody.AcctNo = accnum
	if BrimomDataRequestBody.AcctNo == "" {
=======
	BrimoDataRequestBody.AcctNo = accnum
	if BrimoDataRequestBody.AcctNo == "" {
>>>>>>> 26312a2e32b0152e2754a471b4e1d8b8a9fe0186
		tx.Context.SetTag("desc", utils.MSG_ERROR_PARSING)
		tx.Result = "false"
		c.JSON(http.StatusOK, gin.H{"desc": utils.MSG_ERROR_PARSING})
	} else {
<<<<<<< HEAD
		tx.Context.SetTag("acct_no", BrimomDataRequestBody.AcctNo)

		//Call Data Response From Controller
		message, err := GetBrimoDataPhoenixDB(BrimomDataRequestBody)
=======
		tx.Context.SetTag("acct_no", BrimoDataRequestBody.AcctNo)

		//Call Data Response From Controller
		message, err := GetCeriaDataPhoenixDB(BrimoDataRequestBody)
>>>>>>> 26312a2e32b0152e2754a471b4e1d8b8a9fe0186
		if len(message) == 0 {
			tx.Context.SetTag("desc", utils.MSG_ERROR_NO_DATA)
			tx.Result = "false"
			c.JSON(http.StatusOK, gin.H{"error": utils.MSG_ERROR_NO_DATA})
		} else if err == nil {
			tx.Context.SetTag("desc", utils.MSG_SUCCESS_DATA_FOUND)
			tx.Result = "true"
			GenerateResponse(c, message[0], "200", utils.MSG_SUCCESS_DATA_FOUND)
		} else {
			tx.Context.SetTag("desc", err.Error())
			tx.Result = "false"
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
	}
}

// Ceria Feature Data Hadler, Get Ceria Data, Using Method GET and URL Param
// func CeriaFeatureData(c *gin.Context) {
<<<<<<< HEAD
// 	var BrimoDataRequestBody CeriaDataRequest
=======
// 	var CeriaDataRequestBody CeriaDataRequest
>>>>>>> 26312a2e32b0152e2754a471b4e1d8b8a9fe0186

// 	tx := apm.DefaultTracer.StartTransaction("Get Feature Ceria by Acct No", "request")
// 	defer tx.End()

// 	accnum := c.Param("accnum")

// 	CeriaDataRequestBody.AcctNo = accnum
// 	if CeriaDataRequestBody.AcctNo == "" {
// 		tx.Context.SetTag("desc", utils.MSG_ERROR_PARSING)
// 		tx.Result = "false"
// 		c.JSON(http.StatusOK, gin.H{"desc": utils.MSG_ERROR_PARSING})
// 	} else {
// 		tx.Context.SetTag("acct_no", CeriaDataRequestBody.AcctNo)

// 		message, err := GetCeriaFeatureData(CeriaDataRequestBody)
// 		if len(message) == 0 {
// 			tx.Context.SetTag("desc", utils.MSG_ERROR_NO_DATA)
// 			tx.Result = "false"
// 			c.JSON(http.StatusOK, gin.H{"error": utils.MSG_ERROR_NO_DATA})
// 		} else if err == nil {
// 			tx.Context.SetTag("desc", utils.MSG_SUCCESS_DATA_FOUND)
// 			tx.Result = "true"
// 			c.JSON(http.StatusOK, gin.H{"data": message[0]})
// 		} else {
// 			tx.Context.SetTag("desc", err.Error())
// 			tx.Result = "false"
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 		}
// 	}
// }

// func CeriaDataByEmailV2(c *gin.Context) {
// 	//Call Models Data Request
// 	var CeriaDataByEmailRequestBody CeriaDataByEmailRequest

// 	tx := apm.DefaultTracer.StartTransaction("Get Whitelist Ceria by Email v2", "request")
// 	defer tx.End()

// 	//Read Acctno Request
// 	email := c.Param("email")
// 	fmt.Println(email)

// 	//Append Acctno Request To Models
// 	CeriaDataByEmailRequestBody.Email = email
// 	if CeriaDataByEmailRequestBody.Email == "" {
// 		tx.Context.SetTag("desc", utils.MSG_ERROR_PARSING)
// 		tx.Result = "false"
// 		c.JSON(http.StatusOK, gin.H{"desc": utils.MSG_ERROR_PARSING})
// 	} else {
// 		tx.Context.SetTag("email", CeriaDataByEmailRequestBody.Email)

// 		//Call Data Response From Controller
// 		message, err := GetCeriaDataByEmailPhoenixDBV2(CeriaDataByEmailRequestBody)
// 		if len(message) == 0 {
// 			tx.Context.SetTag("desc", utils.MSG_ERROR_NO_DATA)
// 			tx.Result = "false"
// 			c.JSON(http.StatusOK, gin.H{"error": utils.MSG_ERROR_NO_DATA})
// 		} else if err == nil {
// 			tx.Context.SetTag("desc", utils.MSG_SUCCESS_DATA_FOUND)
// 			tx.Result = "true"
// 			GenerateResponseByEmailV2(c, message[0], "200", utils.MSG_SUCCESS_DATA_FOUND)
// 		} else {
// 			tx.Context.SetTag("desc", err.Error())
// 			tx.Result = "false"
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 		}
// 	}
// }

// func CeriaDataByPhoneV2(c *gin.Context) {
// 	//Call Models Data Request
// 	var CeriaDataByPhoneRequestBody CeriaDataByPhoneRequest

// 	tx := apm.DefaultTracer.StartTransaction("Get Whitelist Ceria by Phone v2", "request")
// 	defer tx.End()

// 	//Read Acctno Request
// 	handphone := c.Param("phone")
// 	fmt.Println(handphone)

// 	//Append Acctno Request To Models
// 	CeriaDataByPhoneRequestBody.Handphone = handphone
// 	if CeriaDataByPhoneRequestBody.Handphone == "" {
// 		tx.Context.SetTag("desc", utils.MSG_ERROR_PARSING)
// 		tx.Result = "false"
// 		c.JSON(http.StatusOK, gin.H{"desc": utils.MSG_ERROR_PARSING})
// 	} else {
// 		tx.Context.SetTag("phone_number", CeriaDataByPhoneRequestBody.Handphone)
// 		//Call Data Response From Controller
// 		message, err := GetCeriaDataByPhonePhoenixDBV2(CeriaDataByPhoneRequestBody)
// 		if len(message) == 0 {
// 			tx.Context.SetTag("desc", utils.MSG_ERROR_NO_DATA)
// 			tx.Result = "false"
// 			c.JSON(http.StatusOK, gin.H{"error": utils.MSG_ERROR_NO_DATA})
// 		} else if err == nil {
// 			tx.Context.SetTag("desc", utils.MSG_SUCCESS_DATA_FOUND)
// 			tx.Result = "true"
// 			GenerateResponseByPhoneV2(c, message[0], "200", utils.MSG_SUCCESS_DATA_FOUND)
// 		} else {
// 			tx.Context.SetTag("desc", err.Error())
// 			tx.Result = "false"
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 		}
// 	}
// }
