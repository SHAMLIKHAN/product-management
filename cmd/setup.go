package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"pm/utils"

	// go package for postgres
	_ "github.com/lib/pq"
)

func validateEnv() error {
	env := []string{
		utils.POSTGRESHOST, utils.POSTGRESPORT, utils.POSTGRESUSER, utils.POSTGRESPASSWORD, utils.POSTGRESDATABASE,
		utils.PORT,
	}
	for _, key := range env {
		if _, ok := os.LookupEnv(key); !ok {
			return fmt.Errorf("App : %s environment variable required but not set", key)
		}
	}
	return nil
}

func prepareDatabase() (*sql.DB, error) {
	db, err := preparePostgres()
	if err != nil {
		return nil, err
	}
	log.Println("App : Database connected successfully!")
	return db, nil
}

func getServerAddr() string {
	port := os.Getenv(utils.PORT)
	addr := ":" + port
	return addr
}
