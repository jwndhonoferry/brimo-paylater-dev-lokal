package main

import (
	"fmt"
	"net/http"

	//"strconv"

	"github.com/gin-gonic/gin"
)

func GenerateResponse(c *gin.Context, req BrimoPaylaterDataResponse, resp_code string, resp_desc string) {
	c.JSON(http.StatusOK, gin.H{
		"AcctNo":             req.AcctNo,
		"Cifno":              req.Cifno,
		"NoOda":              req.NoOda,
		"NamaNasabah":        req.NamaNasabah,
		"UsernameBrimo":      req.UsernameBrimo,
		"NoHandphone":        req.NoHandphone,
		"LimitBrimoPaylater": req.LimitBrimoPaylater,
		"RiskGradeScore":     req.RiskGradeScore,
	})
}

// Message is a utility function to generate the json string message.
// func Message(req CeriaDataResponse) string {
// 	return fmt.Sprintf(
// 		`{ "account_number": %q, "name": %q, "payroll_date": %q, "company_name": %q, "industry_code": %q, "payroll_amount": %q, "employee_contract_start_date": %q, "employee_contract_end_date": %q, "age": %q, "average_saving_amount": %q, "thp_ratio": %q, "pendapatan_netto": %q, "pendidikan_terakhir": %q, "simpanan_britama": %q, "waktu_kepemilikan_simpanan_britama": %q, "inet_banking": %q, "mobile_banking": %q, "primary_phone": %q, "secondary_phone": %q, "status_perkawinan": %q, "freq_mutasi_credit_3": %q, "freq_mutasi_credit_6": %q, "freq_mutasi_credit_12": %q, "freq_mutasi_deb_3": %q, "freq_mutasi_deb_6": %q, "freq_mutasi_deb_12": %q, "total_mutasi_credit_3": %q, "total_mutasi_credit_6": %q, "total_mutasi_credit_12": %q, "total_mutasi_deb_3": %q, "total_mutasi_deb_6": %q, "total_mutasi_deb_12": %q, "rata_mutasi_credit_3": %q, "rata_mutasi_credit_6": %q, "rata_mutasi_credit_12": %q, "rata_mutasi_deb_3": %q, "rata_mutasi_deb_6": %q, "rata_mutasi_deb_12": %q, "handphone": %q, "email": %q }`,
// 		req.AccountNumber,
// 		req.Name,
// 		uint64(req.PayrollDate),
// 		req.CompanyName,
// 		req.IndustryCode,
// 		uint64(req.PayrollAmount),
// 		req.EmployeeContractStartDate,
// 		req.EmployeeContractEndDate,
// 		uint64(req.Age),
// 		uint64(req.AverageSavingAmount),
// 		req.ThpRatio,
// 		uint64(req.PendapatanNetto),
// 		req.PendidikanTerakhir,
// 		req.SimpananBritama,
// 		req.WaktuKepemilikanSimpananBritama,
// 		req.InetBanking,
// 		req.MobileBanking,
// 		req.PrimaryPhone,
// 		req.SecondaryPhone,
// 		req.StatusPerkawinan,
// 		uint64(req.FreqMutasiCredit3),
// 		uint64(req.FreqMutasiCredit6),
// 		uint64(req.FreqMutasiCredit12),
// 		uint64(req.FreqMutasiDeb3),
// 		uint64(req.FreqMutasiDeb6),
// 		uint64(req.FreqMutasiDeb12),
// 		uint64(req.TotalMutasiCredit3),
// 		uint64(req.TotalMutasiCredit6),
// 		uint64(req.TotalMutasiCredit12),
// 		uint64(req.TotalMutasiDeb3),
// 		uint64(req.TotalMutasiDeb6),
// 		uint64(req.TotalMutasiDeb12),
// 		uint64(req.RataMutasiCredit3),
// 		uint64(req.RataMutasiCredit6),
// 		uint64(req.RataMutasiCredit12),
// 		uint64(req.RataMutasiDeb3),
// 		uint64(req.RataMutasiDeb6),
// 		uint64(req.RataMutasiDeb12),
// 		req.Handphone,
// 		req.Email)
// }

// func GenerateResponseByEmailV2(c *gin.Context, req CeriaDataByEmailResponseV2, resp_code string, resp_desc string) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"email":          req.Email,
// 		"account_number": req.AccountNumber,
// 		"handphone":      req.Handphone,
// 	})
// }

// // Message is a utility function to generate the json string message.
// func MessageByEmailV2(req CeriaDataByEmailResponseV2) string {
// 	return fmt.Sprintf(
// 		`{ "email": %q, "account_number": %q, "handphone": %q }`,
// 		req.Email,
// 		req.AccountNumber,
// 		req.Handphone)
// }

// func GenerateResponseByEmail(c *gin.Context, req CeriaDataByEmailResponse, resp_code string, resp_desc string) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"email":                              req.Email,
// 		"account_number":                     req.AccountNumber,
// 		"name":                               req.Name,
// 		"payroll_date":                       req.PayrollDate,
// 		"company_name":                       req.CompanyName,
// 		"industry_code":                      req.IndustryCode,
// 		"payroll_amount":                     req.PayrollAmount,
// 		"employee_contract_start_date":       req.EmployeeContractStartDate,
// 		"employee_contract_end_date":         req.EmployeeContractEndDate,
// 		"age":                                req.Age,
// 		"average_saving_amount":              req.AverageSavingAmount,
// 		"thp_ratio":                          req.ThpRatio,
// 		"pendapatan_netto":                   req.PendapatanNetto,
// 		"pendidikan_terakhir":                req.PendidikanTerakhir,
// 		"simpanan_britama":                   req.SimpananBritama,
// 		"waktu_kepemilikan_simpanan_britama": req.WaktuKepemilikanSimpananBritama,
// 		"inet_banking":                       req.InetBanking,
// 		"mobile_banking":                     req.MobileBanking,
// 		"primary_phone":                      req.PrimaryPhone,
// 		"secondary_phone":                    req.SecondaryPhone,
// 		"status_perkawinan":                  req.StatusPerkawinan,
// 		"freq_mutasi_credit_3":               req.FreqMutasiCredit3,
// 		"freq_mutasi_credit_6":               req.FreqMutasiCredit6,
// 		"freq_mutasi_credit_12":              req.FreqMutasiCredit12,
// 		"freq_mutasi_deb_3":                  req.FreqMutasiDeb3,
// 		"freq_mutasi_deb_6":                  req.FreqMutasiDeb6,
// 		"freq_mutasi_deb_12":                 req.FreqMutasiDeb12,
// 		"total_mutasi_credit_3":              req.TotalMutasiCredit3,
// 		"total_mutasi_credit_6":              req.TotalMutasiCredit6,
// 		"total_mutasi_credit_12":             req.TotalMutasiCredit12,
// 		"total_mutasi_deb_3":                 req.TotalMutasiDeb3,
// 		"total_mutasi_deb_6":                 req.TotalMutasiDeb6,
// 		"total_mutasi_deb_12":                req.TotalMutasiDeb12,
// 		"rata_mutasi_credit_3":               req.RataMutasiCredit3,
// 		"rata_mutasi_credit_6":               req.RataMutasiCredit6,
// 		"rata_mutasi_credit_12":              req.RataMutasiCredit12,
// 		"rata_mutasi_deb_3":                  req.RataMutasiDeb3,
// 		"rata_mutasi_deb_6":                  req.RataMutasiDeb6,
// 		"rata_mutasi_deb_12":                 req.RataMutasiDeb12,
// 		"handphone":                          req.Handphone,
// 	})
// }

// // Message is a utility function to generate the json string message.
// func MessageByEmail(req CeriaDataByEmailResponse) string {
// 	return fmt.Sprintf(
// 		`{ "email": %q, "account_number": %q, "name": %q, "payroll_date": %q, "company_name": %q, "industry_code": %q, "payroll_amount": %q, "employee_contract_start_date": %q, "employee_contract_end_date": %q, "age": %q, "average_saving_amount": %q, "thp_ratio": %q, "pendapatan_netto": %q, "pendidikan_terakhir": %q, "simpanan_britama": %q, "waktu_kepemilikan_simpanan_britama": %q, "inet_banking": %q, "mobile_banking": %q, "primary_phone": %q, "secondary_phone": %q, "status_perkawinan": %q, "freq_mutasi_credit_3": %q, "freq_mutasi_credit_6": %q, "freq_mutasi_credit_12": %q, "freq_mutasi_deb_3": %q, "freq_mutasi_deb_6": %q, "freq_mutasi_deb_12": %q, "total_mutasi_credit_3": %q, "total_mutasi_credit_6": %q, "total_mutasi_credit_12": %q, "total_mutasi_deb_3": %q, "total_mutasi_deb_6": %q, "total_mutasi_deb_12": %q, "rata_mutasi_credit_3": %q, "rata_mutasi_credit_6": %q, "rata_mutasi_credit_12": %q, "rata_mutasi_deb_3": %q, "rata_mutasi_deb_6": %q, "rata_mutasi_deb_12": %q, "handphone": %q }`,
// 		req.Email,
// 		req.AccountNumber,
// 		req.Name,
// 		uint64(req.PayrollDate),
// 		req.CompanyName,
// 		req.IndustryCode,
// 		uint64(req.PayrollAmount),
// 		req.EmployeeContractStartDate,
// 		req.EmployeeContractEndDate,
// 		uint64(req.Age),
// 		uint64(req.AverageSavingAmount),
// 		req.ThpRatio,
// 		uint64(req.PendapatanNetto),
// 		req.PendidikanTerakhir,
// 		req.SimpananBritama,
// 		req.WaktuKepemilikanSimpananBritama,
// 		req.InetBanking,
// 		req.MobileBanking,
// 		req.PrimaryPhone,
// 		req.SecondaryPhone,
// 		req.StatusPerkawinan,
// 		uint64(req.FreqMutasiCredit3),
// 		uint64(req.FreqMutasiCredit6),
// 		uint64(req.FreqMutasiCredit12),
// 		uint64(req.FreqMutasiDeb3),
// 		uint64(req.FreqMutasiDeb6),
// 		uint64(req.FreqMutasiDeb12),
// 		uint64(req.TotalMutasiCredit3),
// 		uint64(req.TotalMutasiCredit6),
// 		uint64(req.TotalMutasiCredit12),
// 		uint64(req.TotalMutasiDeb3),
// 		uint64(req.TotalMutasiDeb6),
// 		uint64(req.TotalMutasiDeb12),
// 		uint64(req.RataMutasiCredit3),
// 		uint64(req.RataMutasiCredit6),
// 		uint64(req.RataMutasiCredit12),
// 		uint64(req.RataMutasiDeb3),
// 		uint64(req.RataMutasiDeb6),
// 		uint64(req.RataMutasiDeb12),
// 		req.Handphone)
// }

// func GenerateResponseByPhoneV2(c *gin.Context, req CeriaDataByPhoneResponseV2, resp_code string, resp_desc string) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"handphone":      req.Handphone,
// 		"account_number": req.AccountNumber,
// 	})
// }

// // Message is a utility function to generate the json string message.
// func MessageByPhoneV2(req CeriaDataByPhoneResponseV2) string {
// 	return fmt.Sprintf(
// 		`{ "handphone": %q, "account_number": %q}`,
// 		req.Handphone,
// 		req.AccountNumber)
// }

// func GenerateResponseByPhone(c *gin.Context, req CeriaDataByPhoneResponse, resp_code string, resp_desc string) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"handphone":                          req.Handphone,
// 		"account_number":                     req.AccountNumber,
// 		"name":                               req.Name,
// 		"payroll_date":                       req.PayrollDate,
// 		"company_name":                       req.CompanyName,
// 		"industry_code":                      req.IndustryCode,
// 		"payroll_amount":                     req.PayrollAmount,
// 		"employee_contract_start_date":       req.EmployeeContractStartDate,
// 		"employee_contract_end_date":         req.EmployeeContractEndDate,
// 		"age":                                req.Age,
// 		"average_saving_amount":              req.AverageSavingAmount,
// 		"thp_ratio":                          req.ThpRatio,
// 		"pendapatan_netto":                   req.PendapatanNetto,
// 		"pendidikan_terakhir":                req.PendidikanTerakhir,
// 		"simpanan_britama":                   req.SimpananBritama,
// 		"waktu_kepemilikan_simpanan_britama": req.WaktuKepemilikanSimpananBritama,
// 		"inet_banking":                       req.InetBanking,
// 		"mobile_banking":                     req.MobileBanking,
// 		"primary_phone":                      req.PrimaryPhone,
// 		"secondary_phone":                    req.SecondaryPhone,
// 		"status_perkawinan":                  req.StatusPerkawinan,
// 		"freq_mutasi_credit_3":               req.FreqMutasiCredit3,
// 		"freq_mutasi_credit_6":               req.FreqMutasiCredit6,
// 		"freq_mutasi_credit_12":              req.FreqMutasiCredit12,
// 		"freq_mutasi_deb_3":                  req.FreqMutasiDeb3,
// 		"freq_mutasi_deb_6":                  req.FreqMutasiDeb6,
// 		"freq_mutasi_deb_12":                 req.FreqMutasiDeb12,
// 		"total_mutasi_credit_3":              req.TotalMutasiCredit3,
// 		"total_mutasi_credit_6":              req.TotalMutasiCredit6,
// 		"total_mutasi_credit_12":             req.TotalMutasiCredit12,
// 		"total_mutasi_deb_3":                 req.TotalMutasiDeb3,
// 		"total_mutasi_deb_6":                 req.TotalMutasiDeb6,
// 		"total_mutasi_deb_12":                req.TotalMutasiDeb12,
// 		"rata_mutasi_credit_3":               req.RataMutasiCredit3,
// 		"rata_mutasi_credit_6":               req.RataMutasiCredit6,
// 		"rata_mutasi_credit_12":              req.RataMutasiCredit12,
// 		"rata_mutasi_deb_3":                  req.RataMutasiDeb3,
// 		"rata_mutasi_deb_6":                  req.RataMutasiDeb6,
// 		"rata_mutasi_deb_12":                 req.RataMutasiDeb12,
// 		"email":                              req.Email,
// 	})
// }

// // Message is a utility function to generate the json string message.
// func MessageByPhone(req CeriaDataByPhoneResponse) string {
// 	return fmt.Sprintf(
// 		`{ "handphone": %q, "account_number": %q, "name": %q, "payroll_date": %q, "company_name": %q, "industry_code": %q, "payroll_amount": %q, "employee_contract_start_date": %q, "employee_contract_end_date": %q, "age": %q, "average_saving_amount": %q, "thp_ratio": %q, "pendapatan_netto": %q, "pendidikan_terakhir": %q, "simpanan_britama": %q, "waktu_kepemilikan_simpanan_britama": %q, "inet_banking": %q, "mobile_banking": %q, "primary_phone": %q, "secondary_phone": %q, "status_perkawinan": %q, "freq_mutasi_credit_3": %q, "freq_mutasi_credit_6": %q, "freq_mutasi_credit_12": %q, "freq_mutasi_deb_3": %q, "freq_mutasi_deb_6": %q, "freq_mutasi_deb_12": %q, "total_mutasi_credit_3": %q, "total_mutasi_credit_6": %q, "total_mutasi_credit_12": %q, "total_mutasi_deb_3": %q, "total_mutasi_deb_6": %q, "total_mutasi_deb_12": %q, "rata_mutasi_credit_3": %q, "rata_mutasi_credit_6": %q, "rata_mutasi_credit_12": %q, "rata_mutasi_deb_3": %q, "rata_mutasi_deb_6": %q, "rata_mutasi_deb_12": %q, "email": %q }`,
// 		req.Handphone,
// 		req.AccountNumber,
// 		req.Name,
// 		uint64(req.PayrollDate),
// 		req.CompanyName,
// 		req.IndustryCode,
// 		uint64(req.PayrollAmount),
// 		req.EmployeeContractStartDate,
// 		req.EmployeeContractEndDate,
// 		uint64(req.Age),
// 		uint64(req.AverageSavingAmount),
// 		req.ThpRatio,
// 		uint64(req.PendapatanNetto),
// 		req.PendidikanTerakhir,
// 		req.SimpananBritama,
// 		req.WaktuKepemilikanSimpananBritama,
// 		req.InetBanking,
// 		req.MobileBanking,
// 		req.PrimaryPhone,
// 		req.SecondaryPhone,
// 		req.StatusPerkawinan,
// 		uint64(req.FreqMutasiCredit3),
// 		uint64(req.FreqMutasiCredit6),
// 		uint64(req.FreqMutasiCredit12),
// 		uint64(req.FreqMutasiDeb3),
// 		uint64(req.FreqMutasiDeb6),
// 		uint64(req.FreqMutasiDeb12),
// 		uint64(req.TotalMutasiCredit3),
// 		uint64(req.TotalMutasiCredit6),
// 		uint64(req.TotalMutasiCredit12),
// 		uint64(req.TotalMutasiDeb3),
// 		uint64(req.TotalMutasiDeb6),
// 		uint64(req.TotalMutasiDeb12),
// 		uint64(req.RataMutasiCredit3),
// 		uint64(req.RataMutasiCredit6),
// 		uint64(req.RataMutasiCredit12),
// 		uint64(req.RataMutasiDeb3),
// 		uint64(req.RataMutasiDeb6),
// 		uint64(req.RataMutasiDeb12),
// 		req.Email)
// }

// =========================================GENERATE DUMMY RESULT=========================================//
func GenerateResponseDummy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"account_number":                     "020601003177535",
		"name":                               "Suci Mailani",
		"payroll_date":                       25,
		"company_name":                       "BANK BRI",
		"industry_code":                      "1",
		"payroll_amount":                     6486632,
		"employee_contract_start_date":       "2012-09-10T00:00:00Z",
		"employee_contract_end_date":         "2044-01-21T00:00:00Z",
		"age":                                0,
		"average_saving_amount":              614766,
		"thp_ratio":                          "",
		"pendapatan_netto":                   6486632,
		"pendidikan_terakhir":                "",
		"simpanan_britama":                   "Y",
		"waktu_kepemilikan_simpanan_britama": "2019-01-29T00:00:00Z",
		"inet_banking":                       "Y",
		"mobile_banking":                     "Y",
		"primary_phone":                      "Y",
		"secondary_phone":                    "N",
		"status_perkawinan":                  "",
		"freq_mutasi_credit_3":               31,
		"freq_mutasi_credit_6":               47,
		"freq_mutasi_credit_12":              0,
		"freq_mutasi_deb_3":                  110,
		"freq_mutasi_deb_6":                  200,
		"freq_mutasi_deb_12":                 0,
		"total_mutasi_credit_3":              18698420,
		"total_mutasi_credit_6":              40906909,
		"total_mutasi_credit_12":             0,
		"total_mutasi_deb_3":                 19940943,
		"total_mutasi_deb_6":                 41536014,
		"total_mutasi_deb_12":                0,
		"rata_mutasi_credit_3":               6232806,
		"rata_mutasi_credit_6":               6817818,
		"rata_mutasi_credit_12":              0,
		"rata_mutasi_deb_3":                  6646981,
		"rata_mutasi_deb_6":                  6922669,
		"rata_mutasi_deb_12":                 0,
		"handphone":                          "6281297620032,081297620032",
		"email":                              "Sucimailani7@gmail.com",
	})
}

// Message is a utility function to generate the json string message.
func MessageDummy() string {
	return fmt.Sprintf(
		`{ "account_number": %q, "name": %q, "payroll_date": %q, "company_name": %q, "industry_code": %q, "payroll_amount": %q, "employee_contract_start_date": %q, "employee_contract_end_date": %q, "age": %q, "average_saving_amount": %q, "thp_ratio": %q, "pendapatan_netto": %q, "pendidikan_terakhir": %q, "simpanan_britama": %q, "waktu_kepemilikan_simpanan_britama": %q, "inet_banking": %q, "mobile_banking": %q, "primary_phone": %q, "secondary_phone": %q, "status_perkawinan": %q, "freq_mutasi_credit_3": %q, "freq_mutasi_credit_6": %q, "freq_mutasi_credit_12": %q, "freq_mutasi_deb_3": %q, "freq_mutasi_deb_6": %q, "freq_mutasi_deb_12": %q, "total_mutasi_credit_3": %q, "total_mutasi_credit_6": %q, "total_mutasi_credit_12": %q, "total_mutasi_deb_3": %q, "total_mutasi_deb_6": %q, "total_mutasi_deb_12": %q, "rata_mutasi_credit_3": %q, "rata_mutasi_credit_6": %q, "rata_mutasi_credit_12": %q, "rata_mutasi_deb_3": %q, "rata_mutasi_deb_6": %q, "rata_mutasi_deb_12": %q, "handphone": %q, "email": %q }`,
		"020601003177535",
		"Suci Mailani",
		25,
		"BANK BRI",
		"1",
		6486632,
		"2012-09-10T00:00:00Z",
		"2044-01-21T00:00:00Z",
		0,
		614766,
		"",
		6486632,
		"",
		"Y",
		"2019-01-29T00:00:00Z",
		"Y",
		"Y",
		"Y",
		"N",
		"",
		31,
		47,
		0,
		110,
		200,
		0,
		18698420,
		40906909,
		0,
		19940943,
		41536014,
		0,
		6232806,
		6817818,
		0,
		6646981,
		6922669,
		0,
		"6281297620032,081297620032",
		"Sucimailani7@gmail.com")
}
