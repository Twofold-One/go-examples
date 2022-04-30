package pack1

import "fmt"

// var will be initialized frist
var firstPack = func() string {
	fmt.Println("var from pack1 initialized")
	return "first"
}

func init() {
	fmt.Println("initializing pack1")
}

func Start() {
	fmt.Println("--start pack1--")
}