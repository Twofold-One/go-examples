package basic

import "fmt"

func MapsExample() {
	// basic map creation
	m := make(map[string]int)
	fmt.Printf("Zero value string-int map: %v\n", m)

	// map literal
	mNums := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	fmt.Println(mNums)

	// map in Go is a hash table, with O(1) for insertion and search
	// Go map is an array of "buckets", which contains max of 8 key/element pairs
	// structure is:
	// key -> hash function -> hash value -> bucket id -> iterate through the keys
	// in bucket and find corresponding element
}