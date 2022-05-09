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
	// createTable(db)
	// createPerson("Maverick", "One", db)
	fmt.Println(getPerson(2, db))
	fmt.Println(getPersons(db))
	updatePersonLastName("Two", 2, db)
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

// read one row
type Person struct {
	id int
	firstname string
	lastname string
}

func getPerson(id int, db *sql.DB) (*Person, error) {
	person := Person{}
	err := db.QueryRow("select id, firstname, lastname from person where id = $1", id).Scan(&person.id, &person.firstname, &person.lastname )
	if err != nil {
		return &person, err
	}
	return &person, nil
}

// read several rows
func getPersons(db *sql.DB) (*[]Person, error) {
	rows, err := db.Query("select id, firstname, lastname from person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	persons := make([]Person, 0)
	for rows.Next() {
		person := Person{}
		if err := rows.Scan(&person.id, &person.firstname, &person.lastname); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return &persons, nil
}

// update row
func updatePersonLastName(lastname string, id int, db *sql.DB) {
	res, err := db.Exec("update person set lastname = $1 where id = $2", lastname, id)
	if err != nil {
		fmt.Println(err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	if affected != 1 {
		fmt.Printf("Something went wrong %d rows were affected expected 1\n", affected)
	} else {
		fmt.Println("Successfully updated")
	}
}
