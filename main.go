package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func main() {
	//connect to database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test user=abdullahi password=")
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to connect to: %v", err.Error()))
	}
	defer conn.Close()

	log.Println("connected to database")

	//test connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot ping database")
	}
	log.Println("pinged database")

	//get rows from database
	err = getAllRows(conn)
	if err != nil {
		log.Fatal("cannot get rows")
	}

	//insert a row
	query := `insert into users (firstname, lastname) values ($1,$2)`
	_, err = conn.Exec(query, "jack", "brown")
	if err != nil {
		log.Fatal("cannot get rows")
	}
	log.Println("inserted a row ")

	//get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal("cannot get rows")
	}

	//update a row
	smt := `update users set firstname = $1 where id = $2 `
	_, err = conn.Exec(smt, "jackie", 5)
	log.Println("updated a row ")

	//get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal("cannot get rows")
	}

	//get one row by id

	//delete a row

	//get rows again
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id,firstname,lastname from users")
	if err != nil {
		log.Fatal("cannot ping database")
	}
	defer rows.Close()

	var firstname, lastname string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstname, &lastname)
		if err != nil {
			log.Fatal("cannot query")
		}
		fmt.Println("record is", id, firstname, lastname)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("error scanning", err)
	}
	fmt.Println("---------------------------------")
	return nil
}
