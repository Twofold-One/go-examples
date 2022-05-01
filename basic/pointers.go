package basic

import "fmt"

// Methods with pointer receivers
type Cat struct {
	Color string
	Age uint8
	Name string
}

// pointer receiver
func (cat *Cat) Rename(newName string) {
	cat.Name = newName
}

// value receiver
func (cat Cat) Renamev2(newName string) {
	cat.Name = newName
}

func PointersExample() {
	i, j := 42, 2701
	
	// regular var call
	fmt.Println("That's the regular var call: ", "i = ",  i, "j = ", j)
	// to get the address of var we use "&" as addres of var
	fmt.Println("Thats's the address of var: &i = ", &i, "&j = ", &j)
	// we also could assign ponter of var to var
	pI := &i
	pJ := &j
	fmt.Println("That's pointers assigned to the variables: pI =", pI, "pJ = ", pJ)
	// we could use "*" operator to get the value of var it points to, it's dereferencing
	fmt.Println("That's the values of var pI and pJ are pointing to: ", *pI, *pJ, "accordingly")
	// when "*" is assigned with a type it's a pointer type
	// var i *int
	// when "*" is assigned with variable name it's dereference
	// *pJ = 555

	cat := Cat{
		Color: "blue",
		Age: 8,
		Name: "Milow",
	}
	cat.Rename("Bob")
	fmt.Println(cat.Name)

	cat.Renamev2("Ben")
	fmt.Println(cat.Name)
}