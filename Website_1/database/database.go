package database

import (
	"database/sql"
	"fmt"
)

//learning ===== to get the value os.Getenv("host"), run 'export host=localhost'
const (
	host     = "hpe-2.c5p6cfuydxbj.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "xz61KLLlvCUJ5Ax2xLag"
	dbname   = "postgres"
)

func StartDatabase() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	return db
}
