package main

import (
	"EffectiveMobile/src/app"
	"EffectiveMobile/src/database"
	_ "EffectiveMobile/src/docs"
)

func main() {
	database.ConnectDB()
	router := app.InitRouter()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
