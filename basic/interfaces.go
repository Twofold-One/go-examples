package basic

import "fmt"

type Human struct {
	Name string
	Age int
	Country string
}

type DomesticAnimal interface {
	ReceiveAffection(from Human)
	GiveAffection(to Human)
}

type iCat struct {
	Name string
}

type iDog struct {
	Name string
}

func (c iCat) ReceiveAffection(from Human) {
	fmt.Printf("The cat named %s has received affection from Human names %s\n", c.Name, from.Name)
}

func (c iCat) GiveAffection(to Human) {
	fmt.Printf("The cat named %s has given affection to Human named %s\n", c.Name, to.Name)
}

func (c iDog) ReceiveAffection(from Human) {
	fmt.Printf("The dog named %s has received affection from Human names %s\n", c.Name, from.Name)
}

func (c iDog) GiveAffection(to Human) {
	fmt.Printf("The dog named %s has given affection to Human named %s\n", c.Name, to.Name)
}

func Pet(animal DomesticAnimal, human Human) {
	animal.GiveAffection(human)
	animal.ReceiveAffection(human)
}

func InterfacesExample() {
	var john Human
	john.Name = "John"

	var c iCat
	c.Name = "Maru"

	var d iDog
	d.Name = "Medor"

	Pet(c, john)
	Pet(d, john)
}