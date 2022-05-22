package appconfiguration

import (
	"flag"
	"fmt"
)

func CommandLineOpts() {
	// flags
	// 1 - flag name; 2 - default value; 3 - help text
	port := flag.Int("port", 4242, "the port on which server will listen")
	flag.Parse()
	fmt.Printf("%d\n", *port)

	var port2 int
	flag.IntVar(&port2, "port2", 4242, "the port on which server will listen")
	flag.Parse()
	fmt.Printf("%d\n", port2)
}
