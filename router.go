package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

// InitRouter is used for initiate router
func InitRouter() *gin.Engine {
	router := gin.New()
	// router := gin.Default()
	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://bri360.bri.co.id", "http://127.0.0.1:8240"},
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST"},
		AllowHeaders:  []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},

		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Use(apmgin.Middleware(router))

	//end point -> call from controller
	router.GET("/api/v1.0/customer/inquiry/:accnum", BrimoData)
	// router.GET("/api/v1.1/customer/inquiry/feature/:accnum", CeriaFeatureData)
	// router.GET("/api/v1.0/customer/inquirybyemail/:email", CeriaDataByEmailV2)
	// router.GET("/api/v1.0/customer/inquirybyphone/:phone", CeriaDataByPhoneV2)
	//router.GET("/api/v1.0/customer/inquirybyemail/:email", CeriaDataByEmail)
	//router.GET("/api/v1.0/customer/inquirybyphone/:phone", CeriaDataByPhone)
	return router
}
