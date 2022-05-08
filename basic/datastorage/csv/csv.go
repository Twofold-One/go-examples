package csv

import (
	"bytes"
	"fmt"
	"os"
)

func CSVStorageExample() {
	// create the file
	f, err := os.OpenFile("./basic/datastorage/csv/test.csv", os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)

	// read the file info
	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info.Mode())

	// change the file mode
	err = f.Chmod(0777)
	if err != nil {
		fmt.Println(err)
	}

	// write to file
	s := [][]string{
		{"age", "gender", "name"},
		{"27", "male", "Curt"},
		{"25", "female", "Grace"},
	}

	// open file or create it
	file, err := os.Create("./basic/datastorage/csv/folks.csv")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var buffer bytes.Buffer
	for _, data := range s {
		buffer.WriteString(fmt.Sprintf("%s,%s,%s\n", data[0], data[1], data[2]))
	}
	n, err := file.Write(buffer.Bytes())
	fmt.Printf("%d bytes written\n", n)
	if err != nil {
		fmt.Println(err)
		return
	}
}