package mapsapp

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Example person list app
// Functionality: load the list of persons from CSV (comma-separated values) file

type person struct {
	firstName string
	lastName string
	gender string
	age string
}

func MapsApp() {
	persons := make(map[string]person)
	order := []string{}

	file, err := os.Open("/home/twofold_one/GitProjects/go/go-examples/basic/mapsapp/persons.csv")
	if err != nil {
		log.Fatalf("impossible to open the file: %s", err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
		
		personID := record[0]
		person := person{
			firstName: record[1],
			lastName: record[2],
			gender: record[3],
			age: record[4],
		}
		order = append(order, personID)
		persons[personID] = person
	}

	// delete first of csv from map
	delete(persons, "personId")

	// retrieve a value from map
	grace := persons["2"]
	fmt.Println(grace)

	// if key doesn't exist we get zero value
	ghost := persons["10"]
	fmt.Println(ghost)

	// two values assignment
	greg, ok := persons["1"]
	if ok {
		fmt.Println(greg)
	} else {
		fmt.Println("No person with id 1")
	}

	ghost2, ok := persons["11"]
	if !ok {
		fmt.Println("No person with id 11")
	} else {
		fmt.Println(ghost2)
	}

	// test presence of a key in map
	if _, ok := persons["20"]; ok {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// delete an entry: func delete(m map[Type]Type, key Type)
	delete(persons, "5")
	fmt.Println(persons)

	// check the length of map
	fmt.Println(len(persons))

	// iterate over map, iteration order != insertation order
	for k, v := range persons {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	// get enties in the insertation order
	for _, k := range order {
		fmt.Printf("key: %v, value: %v\n", k, persons[k])
	}

	// two dimensional map
	map2D := map[int]map[string]string{
		1: {
			"one": "thirst",
		},
		2: {
			"two": "second",
		},
		3: {
			"three": "third",
		},
	}
	fmt.Println(map2D)
}