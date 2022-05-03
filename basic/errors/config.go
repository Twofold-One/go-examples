package errors

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func load() []byte {
	data, err := ioutil.ReadFile("/home/twofold_one/GitProjects/go/go-examples/basic/errors/config.txt")
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func Print() {
	fmt.Println(string(load()))
}

// Custom type that implements the error interface
type HeaderError struct {
	FaultyHeader string
}

func (e *HeaderError) Error() string {
	return fmt.Sprintf("Bad header. Provided %s, expected CONF2", e.FaultyHeader)
}

const fileHeader = "CONF"
const fileHeaderV2 = "CONF2"

// sentinel error
var ErrNoConfigFile = errors.New("no config file at /home/twofold_one/GitProjects/go/go-examples/basic/errors/config.txt")

// alternative Load() function
func Load() (string, error) {
	data, err := ioutil.ReadFile("/home/twofold_one/GitProjects/go/go-examples/basic/errors/config.txt")
	if err != nil {
		// use of sentinel error
		return "", ErrNoConfigFile
	}

	conf := string(data)
	if conf[0:4] != fileHeader {
		// custom error also fmt.Errorf can be used, it will also call errors.New()
		return "", errors.New("the config file header not accepted")
	}

	// if conf[0:4] != fileHeaderV2 {
	// 	return "", &HeaderError{FaultyHeader: conf[0:4]}
	// }

	return conf, nil
}

// wrapping errors using %w
func TransferFileContents(filename string) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("during file transfer impossible to open file: %w", err)
	}
	err = ioutil.WriteFile("/home/twofold_one/GitProjects/go/go-examples/basic/errors/filecontents", contents, 0644)
	if err != nil {
		return fmt.Errorf("during file transfer impossible to write source file: %w", err)
	}
	return nil
}

// custom error type
type ReadingError struct {
	IOError error
	Filename string
}

type WritingError struct {
	IOError error
	Filename string
}

func (e *ReadingError) Error() string {
	return fmt.Sprintf("an error occured while attemting to read the file %s\n", e.Filename)
}

func (e *ReadingError) Unwrap() error {
	return e.IOError
}

func (e *WritingError) Error() string {
	return fmt.Sprintf("an error occured while attemting to write source file %s\n", e.Filename)
}

func (e *WritingError) Unwrap() error {
	return e.IOError
}
