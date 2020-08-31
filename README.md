# product-management

## Setup Postgres Database

    $ cd migration
    $ goose -env development up
    $ cd ..

## Run Development Environment

    $ source config/development.env
    $ go run main.go
