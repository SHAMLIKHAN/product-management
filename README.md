# product-management

## Prerequisites

* Golang v1.14
* Postgres 10.2

## Setup Postgres Database

    $ cd migration
    $ goose -env development up
    $ cd ..

## Run Development Environment

    $ source config/development.env
    $ go run main.go

## Notes

* To run database migration using goose, set database configuration in `/migration/dbconf.yml`
* To setup database, set database environment variables in `/config/development.env`
