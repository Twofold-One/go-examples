package basic

import "fmt"

func TypesExample() {
	// Predeclared types are bool, string, uint (and it's subtypes), int(and it's subtypes)
	// float(and it's subtypes) and complex64/complex128

	// Composite types, type literals

	// array
	var arr [3]uint8

	// slice
	var roomNumbers []uint8

	// pointer
	var myPointer *uint8

	// function
	// var nameDisplayer func(name string) string

	// map
	var score map[string]uint8

	// channel
	var received chan <- bool

	fmt.Println("Composite types:")
	fmt.Println(arr, roomNumbers, myPointer, score, received)

	// new types and structs
	type Country struct {
		Name string
		CapitalCity string
	}

	usa := Country{
		Name: "United States of America",
		CapitalCity: "Washington DC",
	}

	fmt.Printf("It's a country struct: %v; which name is %v and capital is %v\n", usa, usa.Name, usa.CapitalCity)

	// embedded field in the struct
	type Hotel struct {
		Name string
		Rooms uint8
		Country
	}

	hotel := Hotel{
		Name: "Hotel Universe",
		Rooms: 100,
		Country: usa,
	}

	fmt.Printf("It's a hotel struct %v; which located in %v\n", hotel, hotel.Country.Name)
}