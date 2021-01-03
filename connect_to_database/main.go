/*
* This Project indicates the way w Connect to Postgres DB from golang scripts
* First Of All you Need Both Go and postgres server installed previously
* Another Prerequisite is the PostgreSql Driver which is can be installed by
*
        ==>> go get -u github.com/lib/pq
*
*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // _ indicates that don't disturb warning message even if package never used
)

// Next Constants represent the basic information which must be configured before connection
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "a12345"
	dbname   = "postgres"
	search_path  = "DEV"
)



func main() {

	//Next String is one the ways of DB Connection representation
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable search_path=%s", host, port, user, password, dbname, search_path)
	// Second Way
	//connStr := "postgres://postgres:a12345@localhost/postgres?sslmode=disable"

	//The function [sql.Open()] requires you to pass in the name of the driver and the connection string that we set up earlier.
	//We also add some error handling syntax that will deal with any errors we may encounter;
	// for example, we’d get an error if we failed to import the github.com/lib/pq package.
	//db, err := sql.Open("postgres", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// We can also ping our connection which will let us know if our connection is correct or not
	//then we put an error-handling code right after that.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
