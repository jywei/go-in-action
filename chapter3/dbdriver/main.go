package main

import (
	"database/sql"

	_ "github.com/go-in-action/chapter3/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
}
