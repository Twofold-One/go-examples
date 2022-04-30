package methods

import (
	"log"

	"github.com/Twofold-One/go-examples/basic/methods/cart"
)

func MethodsExample() {
	newCart := cart.Cart{}

	totalPrice, err := newCart.TotalPrice()
	if err != nil {
		log.Printf("impossible to compute price of the cart: %s", err)
		return
	}
	log.Printf("Total Price: %v", totalPrice.Display())

	err = newCart.Lock()
	if err != nil {
		log.Printf("impossible to lock the cart: %s", err)
		return
	}
}