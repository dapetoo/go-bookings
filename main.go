package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

func main() {
	//Connect to a Postgres DB
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE"))
	conn, err := sql.Open("pgx", "host=localhost port=5432 user=postgres password=postgres dbname=pg-learn sslmode=disable")
	if err != nil {
		log.Fatal(fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err))
		os.Exit(1)
	}
	defer conn.Close()

	log.Println("Connected to DB successfully")

	err = conn.Ping()
	if err != nil {
		log.Fatal(fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err))
	}

	log.Println("Ping to DB successful")

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("All rows fetched successfully")

	//Insert a row into the DB
	query := `insert into users (first_name, last_name) values ($1, $2)`
	_, err = conn.Exec(query, "Jack", "Nil")
	if err != nil {
		log.Fatal(err)
	}

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a row successfully")

	//Update a row in the DB
	stmt := `update users set first_name = $1 where id = $2`
	_, err = conn.Exec(stmt, "Jackie", 4)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated a row successfully")

	//Get Row By ID
	query = `select id, first_name, last_name from users where id = $1`

	var firstName, lastName string
	var id int
	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Query returns ", id, firstName, lastName)

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	//Delete a row from the DB
	query = `delete from users where id = $1`
	_, err = conn.Exec(query, 4)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted a row successfully")

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
}

// Get all Rows
func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, first_name, last_name from users")

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("All rows fetched successfully, and they are", "ID: ", id, "First Name: ", firstName, "Last Name: ", lastName)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("=====================================")

	return nil
}
