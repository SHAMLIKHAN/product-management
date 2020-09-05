package cmd

import (
	"log"
)

// Begin : Beginning of the app
func Begin() {
	err := validateEnv()
	if err != nil {
		log.Println("App : Environment configuration failed!")
		panic(err)
	}
	db, err := prepareDatabase()
	if err != nil {
		log.Println("App : Database connection failed!")
		panic(err)
	} else {
		app := NewApp(db)
		app.Serve(getServerAddr())
	}
}
