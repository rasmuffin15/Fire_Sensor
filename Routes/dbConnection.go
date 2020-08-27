package Routes

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
//Postgres DB details
//Currently local
const (
	host     = "localhost"
	port     = 5432
	user     = "thomasrasmussen"
	password = "Colorado@Boulder15"
	dbname   = "firesensor"
)

//Connect to postgres database
//Returns connection
func DBConn() *sql.DB {
	//Holds database details
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Creates connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}

//Close connection to database
func DBClose(db *sql.DB) {
	db.Close()
	fmt.Println("Closed Database")
}