package appconfiguration

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// To create env var named MYVAR type following command in terminal:
// $ export MYVAR=value

// To env vars to app:
// $ export MYVAR=test && export MYVAR2=test2 && ./myApp

func EnvVars() {
	// to retreive env vars
	envVar1 := os.Getenv("MYVAR1")
	envVar2 := os.Getenv("MYVAR2")
	fmt.Printf("myVar1: %s\n", envVar1)
	fmt.Printf("myVar2: %s\n", envVar2)

	// to retreive env vars using LookupEnv
	port, found := os.LookupEnv("DB_PORT")
	if !found {
		log.Fatal("impossible to start up, DB_PORT env var is needed")
	}
	portParsed, err := strconv.ParseUint(port, 10, 32)
	if err != nil {
		log.Fatalf("impossible to parse db port: %s", err)
	}
	log.Println(portParsed)

	// to get all env vars
	fmt.Println(os.Environ())

	// to set env var
	// 1 - name of var; 2 - it's value
	err = os.Setenv("MYVAR3", "test3")
	if err != nil {
		panic(err)
	}
}
