package main

type BrimoPaylaterDataByAcctNoRequest struct {
	AcctNo string `json:acctno`
}

type BrimoPaylaterDataByCifnoRequest struct {
	Cifno string `json:cifno`
}

type BrimoPaylaterDataResponse struct {
	AcctNo             string  `json:"acctno" gorm:"column:acctno"`
	Cifno              string  `json:"cifno" gorm:"column:cifno"`
	NoOda              string  `json:"no_oda" gorm:"column:no_oda"`
	NamaNasabah        string  `json:"nama_nasabah" gorm:"column:nama_nasabah"`
	UsernameBrimo      string  `json:"username_brimo" gorm:"column:username_brimo"`
	NoHandphone        string  `json:"no_handphone" gorm:"column:no_handphone"`
	LimitBrimoPaylater float64 `json:"limit_brimo_paylater" gorm:"column:limit_brimo_paylater"`
	RiskGradeScore     uint64  `json:"risk_grade_score" gorm:"column:risk_grade_score"`
}

type BrimoPaylaterDataCifnoResponse struct {
	AcctNo             string  `json:"acctno" gorm:"column:acctno"`
	Cifno              string  `json:"cifno" gorm:"column:cifno"`
	NoOda              string  `json:"no_oda" gorm:"column:no_oda"`
	NamaNasabah        string  `json:"nama_nasabah" gorm:"column:nama_nasabah"`
	UsernameBrimo      string  `json:"username_brimo" gorm:"column:username_brimo"`
	NoHandphone        string  `json:"no_handphone" gorm:"column:no_handphone"`
	LimitBrimoPaylater float64 `json:"limit_brimo_paylater" gorm:"column:limit_brimo_paylater"`
	RiskGradeScore     uint64  `json:"risk_grade_score" gorm:"column:risk_grade_score"`
}

// CeriaDataRequest is request for ceria data service
// type CeriaDataRequest struct {
// 	AcctNo string `json:accnum`
// }

// // CeriaDataResponse is result for ceria data service
// type CeriaDataResponse struct {
// 	AccountNumber                   string  `json:"account_number" gorm:"column:ACCOUNT_NUMBER"`
// 	Name                            string  `json:"name" gorm:"column:NAME"`
// 	PayrollDate                     uint64  `json:"payroll_date" gorm:"column:PAYROLL_DATE"`
// 	CompanyName                     string  `json:"company_name" gorm:"column:COMPANY_NAME"`
// 	IndustryCode                    string  `json:"industry_code" gorm:"column:INDUSTRY_CODE"`
// 	PayrollAmount                   float64 `json:"payroll_amount" gorm:"column:PAYROLL_AMOUNT"`
// 	EmployeeContractStartDate       string  `json:"employee_contract_start_date" gorm:"column:EMPLOYEE_CONTRACT_START_DATE"`
// 	EmployeeContractEndDate         string  `json:"employee_contract_end_date" gorm:"column:EMPLOYEE_CONTRACT_END_DATE"`
// 	Age                             uint64  `json:"age" gorm:"column:AGE"`
// 	AverageSavingAmount             float64 `json:"average_saving_amount" gorm:"column:AVERAGE_SAVING_AMOUNT"`
// 	ThpRatio                        string  `json:"thp_ratio" gorm:"column:THP_RATIO"`
// 	PendapatanNetto                 float64 `json:"pendapatan_netto" gorm:"column:PENDAPATAN_NETTO"`
// 	PendidikanTerakhir              string  `json:"pendidikan_terakhir" gorm:"column:PENDIDIKAN_TERAKHIR"`
// 	SimpananBritama                 string  `json:"simpanan_britama" gorm:"column:SIMPANAN_BRITAMA"`
// 	WaktuKepemilikanSimpananBritama string  `json:"waktu_kepemilikan_simpanan_britama" gorm:"column:WAKTU_KEPEMILIKAN_SIMPANAN_BRITAMA"`
// 	InetBanking                     string  `json:"inet_banking" gorm:"column:INET_BANKING"`
// 	MobileBanking                   string  `json:"mobile_banking" gorm:"column:MOBILE_BANKING"`
// 	PrimaryPhone                    string  `json:"primary_phone" gorm:"column:PRIMARY_PHONE"`
// 	SecondaryPhone                  string  `json:"secondary_phone" gorm:"column:SECONDARY_PHONE"`
// 	StatusPerkawinan                string  `json:"status_perkawinan" gorm:"column:STATUS_PERKAWINAN"`
// 	FreqMutasiCredit3               float64 `json:"freq_mutasi_credit_3" gorm:"column:FREQ_MUTASI_CREDIT_3"`
// 	FreqMutasiCredit6               float64 `json:"freq_mutasi_credit_6" gorm:"column:FREQ_MUTASI_CREDIT_6"`
// 	FreqMutasiCredit12              float64 `json:"freq_mutasi_credit_12" gorm:"column:FREQ_MUTASI_CREDIT_12"`
// 	FreqMutasiDeb3                  float64 `json:"freq_mutasi_deb_3" gorm:"column:FREQ_MUTASI_DEB_3"`
// 	FreqMutasiDeb6                  float64 `json:"freq_mutasi_deb_6" gorm:"column:FREQ_MUTASI_DEB_6"`
// 	FreqMutasiDeb12                 float64 `json:"freq_mutasi_deb_12" gorm:"column:FREQ_MUTASI_DEB_12"`
// 	TotalMutasiCredit3              float64 `json:"total_mutasi_credit_3" gorm:"column:TOTAL_MUTASI_CREDIT_3"`
// 	TotalMutasiCredit6              float64 `json:"total_mutasi_credit_6" gorm:"column:TOTAL_MUTASI_CREDIT_6"`
// 	TotalMutasiCredit12             float64 `json:"total_mutasi_credit_12" gorm:"column:TOTAL_MUTASI_CREDIT_12"`
// 	TotalMutasiDeb3                 float64 `json:"total_mutasi_deb_3" gorm:"column:TOTAL_MUTASI_DEB_3"`
// 	TotalMutasiDeb6                 float64 `json:"total_mutasi_deb_6" gorm:"column:TOTAL_MUTASI_DEB_6"`
// 	TotalMutasiDeb12                float64 `json:"total_mutasi_deb_12" gorm:"column:TOTAL_MUTASI_DEB_12"`
// 	RataMutasiCredit3               float64 `json:"rata_mutasi_credit_3" gorm:"column:RATA_MUTASI_CREDIT_3"`
// 	RataMutasiCredit6               float64 `json:"rata_mutasi_credit_6" gorm:"column:RATA_MUTASI_CREDIT_6"`
// 	RataMutasiCredit12              float64 `json:"rata_mutasi_credit_12" gorm:"column:RATA_MUTASI_CREDIT_12"`
// 	RataMutasiDeb3                  float64 `json:"rata_mutasi_deb_3" gorm:"column:RATA_MUTASI_DEB_3"`
// 	RataMutasiDeb6                  float64 `json:"rata_mutasi_deb_6" gorm:"column:RATA_MUTASI_DEB_6"`
// 	RataMutasiDeb12                 float64 `json:"rata_mutasi_deb_12" gorm:"column:RATA_MUTASI_DEB_12"`
// 	Handphone                       string  `json:"handphone" gorm:"column:HANDPHONE"`
// 	Email                           string  `json:"email" gorm:"column:EMAIL"`
// 	RespCode                        string  `json:"resp_code"`
// 	RespDesc                        string  `json:"resp_desc"`
// 	FlagPekerjaBri                  string  `json:"flag_pekerja_bri" gorm:"column:FLAG_PEKERJA_BRI"`
// 	BypassValidation                string  `json:"bypass_validation" gorm:"column:BYPASS_VALIDATION"`
// 	Pg                              string  `json:"pg" gorm:"column:PG"`
// 	FlagOverride                    uint64  `json:"flag_override" gorm:"column:FLAG_OVERRIDE"`
// 	LimitOverride                   float64 `json:"limit_override" gorm:"column:LIMIT_OVERRIDE"`
// }

// type CeriaFeatureDataResponse struct {
// 	Acctno                          float64 `json:"acctno" gorm:"column:ACCTNO"`
// 	WaktuKepemilikanSimpananBritama string  `json:"waktu_kepemilikan_simpanan_britama" gorm:"column:WAKTU_KEPEMILIKAN_SIMPANAN_BRITAMA"`
// 	Cifno                           string  `json:"cifno" gorm:"column:CIFNO"`
// 	TanggalLahir                    string  `json:"tanggal_lahir" gorm:"column:TANGGAL_LAHIR"`
// 	AverageSavingAmount             float64 `json:"average_saving_amount" gorm:"column:AVERAGE_SAVING_AMOUNT"`
// 	FreqMutasiCredit3               float64 `json:"freq_mutasi_credit_3" gorm:"column:FREQ_MUTASI_CREDIT_3"`
// 	FreqMutasiDeb3                  float64 `json:"freq_mutasi_deb_3" gorm:"column:FREQ_MUTASI_DEB_3"`
// 	TotalMutasiCredit3              float64 `json:"total_mutasi_credit_3" gorm:"column:TOTAL_MUTASI_CREDIT_3"`
// 	FlagIncome                      int     `json:"flag_income" gorm:"column:FLAG_INCOME"`
// 	IncomeCalculation               float64 `json:"income_calculation" gorm:"column:INCOME_CALCULATION"`
// 	PredProba                       float32 `json:"pred_proba" gorm:"column:PRED_PROBA"`
// 	ModifiedDate                    string  `json:"modified_data" gorm:"column:MODIFIED_DATE"`
// 	Ds                              string  `json:"ds" gorm:"column:DS"`
// }

// type CeriaDataByEmailRequest struct {
// 	Email string `json:"email"`
// }
// type CeriaDataByEmailResponseV2 struct {
// 	Email         string `json:"email" gorm:"column:EMAIL"`
// 	AccountNumber string `json:"account_number" gorm:"column:ACCOUNT_NUMBER"`
// 	Handphone     string `json:"handphone" gorm:"column:HANDPHONE"`
// }
// type CeriaDataByEmailResponse struct {
// 	Email                           string  `json:"email" gorm:"column:EMAIL"`
// 	AccountNumber                   string  `json:"account_number" gorm:"column:ACCOUNT_NUMBER"`
// 	Name                            string  `json:"name" gorm:"column:NAME"`
// 	PayrollDate                     uint64  `json:"payroll_date" gorm:"column:PAYROLL_DATE"`
// 	CompanyName                     string  `json:"company_name" gorm:"column:COMPANY_NAME"`
// 	IndustryCode                    string  `json:"industry_code" gorm:"column:INDUSTRY_CODE"`
// 	PayrollAmount                   float64 `json:"payroll_amount" gorm:"column:PAYROLL_AMOUNT"`
// 	EmployeeContractStartDate       string  `json:"employee_contract_start_date" gorm:"column:EMPLOYEE_CONTRACT_START_DATE"`
// 	EmployeeContractEndDate         string  `json:"employee_contract_end_date" gorm:"column:EMPLOYEE_CONTRACT_END_DATE"`
// 	Age                             uint64  `json:"age" gorm:"column:AGE"`
// 	AverageSavingAmount             float64 `json:"average_saving_amount" gorm:"column:AVERAGE_SAVING_AMOUNT"`
// 	ThpRatio                        string  `json:"thp_ratio" gorm:"column:THP_RATIO"`
// 	PendapatanNetto                 float64 `json:"pendapatan_netto" gorm:"column:PENDAPATAN_NETTO"`
// 	PendidikanTerakhir              string  `json:"pendidikan_terakhir" gorm:"column:PENDIDIKAN_TERAKHIR"`
// 	SimpananBritama                 string  `json:"simpanan_britama" gorm:"column:SIMPANAN_BRITAMA"`
// 	WaktuKepemilikanSimpananBritama string  `json:"waktu_kepemilikan_simpanan_britama" gorm:"column:WAKTU_KEPEMILIKAN_SIMPANAN_BRITAMA"`
// 	InetBanking                     string  `json:"inet_banking" gorm:"column:INET_BANKING"`
// 	MobileBanking                   string  `json:"mobile_banking" gorm:"column:MOBILE_BANKING"`
// 	PrimaryPhone                    string  `json:"primary_phone" gorm:"column:PRIMARY_PHONE"`
// 	SecondaryPhone                  string  `json:"secondary_phone" gorm:"column:SECONDARY_PHONE"`
// 	StatusPerkawinan                string  `json:"status_perkawinan" gorm:"column:STATUS_PERKAWINAN"`
// 	FreqMutasiCredit3               float64 `json:"freq_mutasi_credit_3" gorm:"column:FREQ_MUTASI_CREDIT_3"`
// 	FreqMutasiCredit6               float64 `json:"freq_mutasi_credit_6" gorm:"column:FREQ_MUTASI_CREDIT_6"`
// 	FreqMutasiCredit12              float64 `json:"freq_mutasi_credit_12" gorm:"column:FREQ_MUTASI_CREDIT_12"`
// 	FreqMutasiDeb3                  float64 `json:"freq_mutasi_deb_3" gorm:"column:FREQ_MUTASI_DEB_3"`
// 	FreqMutasiDeb6                  float64 `json:"freq_mutasi_deb_6" gorm:"column:FREQ_MUTASI_DEB_6"`
// 	FreqMutasiDeb12                 float64 `json:"freq_mutasi_deb_12" gorm:"column:FREQ_MUTASI_DEB_12"`
// 	TotalMutasiCredit3              float64 `json:"total_mutasi_credit_3" gorm:"column:TOTAL_MUTASI_CREDIT_3"`
// 	TotalMutasiCredit6              float64 `json:"total_mutasi_credit_6" gorm:"column:TOTAL_MUTASI_CREDIT_6"`
// 	TotalMutasiCredit12             float64 `json:"total_mutasi_credit_12" gorm:"column:TOTAL_MUTASI_CREDIT_12"`
// 	TotalMutasiDeb3                 float64 `json:"total_mutasi_deb_3" gorm:"column:TOTAL_MUTASI_DEB_3"`
// 	TotalMutasiDeb6                 float64 `json:"total_mutasi_deb_6" gorm:"column:TOTAL_MUTASI_DEB_6"`
// 	TotalMutasiDeb12                float64 `json:"total_mutasi_deb_12" gorm:"column:TOTAL_MUTASI_DEB_12"`
// 	RataMutasiCredit3               float64 `json:"rata_mutasi_credit_3" gorm:"column:RATA_MUTASI_CREDIT_3"`
// 	RataMutasiCredit6               float64 `json:"rata_mutasi_credit_6" gorm:"column:RATA_MUTASI_CREDIT_6"`
// 	RataMutasiCredit12              float64 `json:"rata_mutasi_credit_12" gorm:"column:RATA_MUTASI_CREDIT_12"`
// 	RataMutasiDeb3                  float64 `json:"rata_mutasi_deb_3" gorm:"column:RATA_MUTASI_DEB_3"`
// 	RataMutasiDeb6                  float64 `json:"rata_mutasi_deb_6" gorm:"column:RATA_MUTASI_DEB_6"`
// 	RataMutasiDeb12                 float64 `json:"rata_mutasi_deb_12" gorm:"column:RATA_MUTASI_DEB_12"`
// 	Handphone                       string  `json:"handphone" gorm:"column:HANDPHONE"`
// 	RespCode                        string  `json:"resp_code"`
// 	RespDesc                        string  `json:"resp_desc"`
// }

// type CeriaDataByPhoneRequest struct {
// 	Handphone string `json:handphone`
// }
// type CeriaDataByPhoneResponseV2 struct {
// 	Handphone     string `json:"handphone" gorm:"column:HANDPHONE"`
// 	AccountNumber string `json:"account_number" gorm:"column:ACCOUNT_NUMBER"`
// }
// type CeriaDataByPhoneResponse struct {
// 	Handphone                       string  `json:"handphone" gorm:"column:HANDPHONE"`
// 	AccountNumber                   string  `json:"account_number" gorm:"column:ACCOUNT_NUMBER"`
// 	Name                            string  `json:"name" gorm:"column:NAME"`
// 	PayrollDate                     uint64  `json:"payroll_date" gorm:"column:PAYROLL_DATE"`
// 	CompanyName                     string  `json:"company_name" gorm:"column:COMPANY_NAME"`
// 	IndustryCode                    string  `json:"industry_code" gorm:"column:INDUSTRY_CODE"`
// 	PayrollAmount                   float64 `json:"payroll_amount" gorm:"column:PAYROLL_AMOUNT"`
// 	EmployeeContractStartDate       string  `json:"employee_contract_start_date" gorm:"column:EMPLOYEE_CONTRACT_START_DATE"`
// 	EmployeeContractEndDate         string  `json:"employee_contract_end_date" gorm:"column:EMPLOYEE_CONTRACT_END_DATE"`
// 	Age                             uint64  `json:"age" gorm:"column:AGE"`
// 	AverageSavingAmount             float64 `json:"average_saving_amount" gorm:"column:AVERAGE_SAVING_AMOUNT"`
// 	ThpRatio                        string  `json:"thp_ratio" gorm:"column:THP_RATIO"`
// 	PendapatanNetto                 float64 `json:"pendapatan_netto" gorm:"column:PENDAPATAN_NETTO"`
// 	PendidikanTerakhir              string  `json:"pendidikan_terakhir" gorm:"column:PENDIDIKAN_TERAKHIR"`
// 	SimpananBritama                 string  `json:"simpanan_britama" gorm:"column:SIMPANAN_BRITAMA"`
// 	WaktuKepemilikanSimpananBritama string  `json:"waktu_kepemilikan_simpanan_britama" gorm:"column:WAKTU_KEPEMILIKAN_SIMPANAN_BRITAMA"`
// 	InetBanking                     string  `json:"inet_banking" gorm:"column:INET_BANKING"`
// 	MobileBanking                   string  `json:"mobile_banking" gorm:"column:MOBILE_BANKING"`
// 	PrimaryPhone                    string  `json:"primary_phone" gorm:"column:PRIMARY_PHONE"`
// 	SecondaryPhone                  string  `json:"secondary_phone" gorm:"column:SECONDARY_PHONE"`
// 	StatusPerkawinan                string  `json:"status_perkawinan" gorm:"column:STATUS_PERKAWINAN"`
// 	FreqMutasiCredit3               float64 `json:"freq_mutasi_credit_3" gorm:"column:FREQ_MUTASI_CREDIT_3"`
// 	FreqMutasiCredit6               float64 `json:"freq_mutasi_credit_6" gorm:"column:FREQ_MUTASI_CREDIT_6"`
// 	FreqMutasiCredit12              float64 `json:"freq_mutasi_credit_12" gorm:"column:FREQ_MUTASI_CREDIT_12"`
// 	FreqMutasiDeb3                  float64 `json:"freq_mutasi_deb_3" gorm:"column:FREQ_MUTASI_DEB_3"`
// 	FreqMutasiDeb6                  float64 `json:"freq_mutasi_deb_6" gorm:"column:FREQ_MUTASI_DEB_6"`
// 	FreqMutasiDeb12                 float64 `json:"freq_mutasi_deb_12" gorm:"column:FREQ_MUTASI_DEB_12"`
// 	TotalMutasiCredit3              float64 `json:"total_mutasi_credit_3" gorm:"column:TOTAL_MUTASI_CREDIT_3"`
// 	TotalMutasiCredit6              float64 `json:"total_mutasi_credit_6" gorm:"column:TOTAL_MUTASI_CREDIT_6"`
// 	TotalMutasiCredit12             float64 `json:"total_mutasi_credit_12" gorm:"column:TOTAL_MUTASI_CREDIT_12"`
// 	TotalMutasiDeb3                 float64 `json:"total_mutasi_deb_3" gorm:"column:TOTAL_MUTASI_DEB_3"`
// 	TotalMutasiDeb6                 float64 `json:"total_mutasi_deb_6" gorm:"column:TOTAL_MUTASI_DEB_6"`
// 	TotalMutasiDeb12                float64 `json:"total_mutasi_deb_12" gorm:"column:TOTAL_MUTASI_DEB_12"`
// 	RataMutasiCredit3               float64 `json:"rata_mutasi_credit_3" gorm:"column:RATA_MUTASI_CREDIT_3"`
// 	RataMutasiCredit6               float64 `json:"rata_mutasi_credit_6" gorm:"column:RATA_MUTASI_CREDIT_6"`
// 	RataMutasiCredit12              float64 `json:"rata_mutasi_credit_12" gorm:"column:RATA_MUTASI_CREDIT_12"`
// 	RataMutasiDeb3                  float64 `json:"rata_mutasi_deb_3" gorm:"column:RATA_MUTASI_DEB_3"`
// 	RataMutasiDeb6                  float64 `json:"rata_mutasi_deb_6" gorm:"column:RATA_MUTASI_DEB_6"`
// 	RataMutasiDeb12                 float64 `json:"rata_mutasi_deb_12" gorm:"column:RATA_MUTASI_DEB_12"`
// 	RespCode                        string  `json:"resp_code"`
// 	RespDesc                        string  `json:"resp_desc"`
// 	Email                           string  `json:"email" gorm:"column:EMAIL"`
// }

// // ============================================ OTHERS ============================================//
type MasterAcctno struct {
	AccountNumber string `json:"account_number" gorm:"column:ACCOUNT_NUMBER"`
	Cifno         string `json:"cifno" gorm:"column:CIFNO"`
	IdNumber      string `json:"id_number" gorm:"column:ID_NUMBER"`
}
