package pack2

import "fmt"

// var will be initialized frist
var secondPack = func() string {
	fmt.Println("var from pack2 initialized")
	return "second"
}

func init() {
	fmt.Println("initializing pack2")
}

func Start() {
	fmt.Println("--start pack2--")
}