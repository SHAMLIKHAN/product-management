package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"pm/utils"
)

func preparePostgres() (*sql.DB, error) {
	url, err := getDatabaseURL()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDatabaseURL() (string, error) {
	env, err := getEnv()
	if err != nil {
		return utils.EmptyString, err
	}
	psql := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		env[utils.POSTGRESHOST],
		env[utils.POSTGRESPORT],
		env[utils.POSTGRESUSER],
		env[utils.POSTGRESPASSWORD],
		env[utils.POSTGRESDATABASE],
	)
	return psql, nil
}

func getEnv() (map[string]string, error) {
	env := make(map[string]string)
	host := os.Getenv(utils.POSTGRESHOST)
	port := os.Getenv(utils.POSTGRESPORT)
	user := os.Getenv(utils.POSTGRESUSER)
	password := os.Getenv(utils.POSTGRESPASSWORD)
	database := os.Getenv(utils.POSTGRESDATABASE)
	env[utils.POSTGRESHOST] = host
	env[utils.POSTGRESPORT] = port
	env[utils.POSTGRESUSER] = user
	env[utils.POSTGRESPASSWORD] = password
	env[utils.POSTGRESDATABASE] = database
	return env, nil
}
