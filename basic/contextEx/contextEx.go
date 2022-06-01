package contextEx

import (
	"context"
	"fmt"
	"log"
	"time"
)

func ContextEx() {
	// create Background concept / root ocontext
	ctx := context.Background()
	foo1(ctx, 1)
	go ContextExServer()
	log.Println("server started")
	time.Sleep(time.Millisecond * 500)
	ContextExClient()
}

// Go idioms:
// The context is the first argument of a function
// The context argument named "ctx"
func foo1(ctx context.Context, a int) {
	fmt.Println(ctx)
	fmt.Println(a)
}
