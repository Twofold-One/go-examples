package errors

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func ErrorsExample() {
	Print()
	
	confData, err := Load()
	if err != nil {
		log.Fatalf("Imposible to load config.txt: %s", err)
	}
	fmt.Println(confData)

	// os.Exit() with exit code
	// exit code 0 means there is no issues
	// exit code 1 to 125 means there was an issue
	// os.Exit(0)

	// wrapping errors
	err = TransferFileContents("/home/twofold_one/GitProjects/go/go-examples/basic/errors/file.txt")
	if err != nil {
		log.Printf("error occured: %s", err)
	}

	// detect a sentinel error with errors.Is(): func Is(err, target error) bool
	readCSV()
}

func readCSV() {
	file, err := os.Open("/home/twofold_one/GitProjects/go/go-examples/basic/errors/persons.csv")
	defer file.Close()
	if err != nil {
		log.Printf("impossible to open file %s", err)
		return
	}

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}

	transferFileWrapperFunc()
}

func transferFileWrapperFunc() {
	var readingError *ReadingError
	var writingError *WritingError
	
	err := TransferFileContents("/home/twofold_one/GitProjects/go/go-examples/basic/errors/file.txt")
	if errors.As(err, &readingError) {
		log.Fatalf("error of reading occured: %s: %s", readingError, readingError.Unwrap())
	}
	if errors.As(err, &writingError) {
		log.Fatalf("error of writing occured: %s: %s", writingError, writingError.Unwrap())
	}
}