package main

import (
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	// InitConfig() //Register Connection Addres Here
	InitDb() //Open And Close Connection Once Time Here
	router := InitRouter()
	router.Run(":9192")
}
