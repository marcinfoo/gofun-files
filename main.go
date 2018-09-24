package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func writeFile(file *string, value string) error {

	if _, err := os.Stat(*file); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fmt.Printf("File [{%s}] does not exist. Will be created.", *file)
	}

	fmt.Printf("Writing to file [%s]\n%s\n", *file, value)
	err := ioutil.WriteFile(*file, []byte(value), 0644)

	return err

}

func readFile(file *string) ([]byte, error) {

	// check if file exists
	if _, err := os.Stat(*file); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fmt.Printf("File [{%s}] does not exist.", *file)
		return make([]byte, 0), errors.New("file does not exist")
	}

	fmt.Printf("Reading from file [%s]\n", *file)
	in, err := ioutil.ReadFile(*file)

	return in, err
}

func main() {

	var fileName = flag.String("fileName", "file.txt", "file to be used in read/write")

	flag.Parse()

	// write current timestamp to file
	timeNow := time.Now()
	timeAsStr := fmt.Sprintf("%s", timeNow.UTC())

	err := writeFile(fileName, timeAsStr)

	if err != nil {
		log.Fatalf("something went serously wrong")
	}

	in, err := readFile(fileName)

	fmt.Println(string(in))
}
