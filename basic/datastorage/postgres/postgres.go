package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func PostgresExample() {
	// connecting to db using pgx driver to local db container
	db, err := sql.Open("pgx", "postgres://root:secret@localhost:5434/people")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	createTable(db)
	createPerson("Maverick", "One", db)
}

func createTable(db *sql.DB) {
		// load sql script
		f, err := os.Open("./basic/datastorage/postgres/sql/table_creation.sql")
		if err != nil {
			fmt.Println(err)
			return
		}
		// slice of bytes of sql query
		b, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		// execute the script
		res, err := db.Exec(string(b))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
}

// creating new person
func createPerson(firstname string, lastname string, db *sql.DB) (int, error) {
	insertedId := 0
	err := db.QueryRow("insert into person (create_time, firstname, lastname) values (now(), $1, $2) returning id;", firstname, lastname).Scan(&insertedId)
	if err != nil {
		return 0, err
	}
	if insertedId == 0 {
		return 0, errors.New("something went wrong, inserted id is equl to zero")
	}
	return insertedId, nil
}

// TODO