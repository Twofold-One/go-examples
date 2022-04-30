package packages

import (
	"fmt"

	"github.com/Twofold-One/go-examples/basic/packages/pack1"
	"github.com/Twofold-One/go-examples/basic/packages/pack2"
)

func init() {
	fmt.Println("initializing main package example")
}

func PackagesExample() {
	pack1.Start()
	pack2.Start()
	fmt.Println("--start main package example--")
}