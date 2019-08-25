package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

//ConnectDB - Method to connect to the database
func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("APP_DB_URL"))
	if err != nil {
		log.Fatalln("Error while connecting to DB")
		os.Exit(1)
	}

	//to open a connection -> db is the global variable
	db, err = sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}

	// ping db to test
	err = db.Ping() //if ping does not retunr anything that means connection is established successfully
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return db
}
