package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const path string = "reader_interface/demo_txt.txt"

func main() {

	x, _ := simpleRead(path)

	fmt.Printf("simple read results: %v", x)

	y, _ := robustRead(path)

	fmt.Printf("robust read results: %v", y)
}

// simple implementation -- assumes the entire content of the file
// can be read into the buffer
func simpleRead(path string) (string, error) {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal("couldn't open file")
	}

	defer f.Close()

	l := 10

	buf := make([]byte, l)

	// fmt.Printf("Contents of buf before reading file: %v\n", buf)

	n, err := f.Read(buf)

	if err != nil {
		return "", err
	}

	// fmt.Printf("Contents of buf after reading file: %v\n", buf)

	j := buf[:n]

	// fmt.Printf("Contents of buf through n bytes: %v\n", j)

	k := string(j)

	// fmt.Printf("Contents of buf through n bytes as string: %v\n", k)

	return k, nil

}

// more robust implementation
// appends content read into the buffer into a variable
// this will work regardless of the size of the buffer and the file
func robustRead(path string) (string, error) {
	f, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer f.Close()

	//prepare buffer and content
	buf := make([]byte, 10)
	var content []byte

	//iteratively read until end of file
	for {
		n, err := f.Read(buf)

		if n > 0 {
			//append freshly-read bytes to our content variable
			content = append(content, buf[:n]...)
		}
		if err == io.EOF {
			//end reading bc we've reached the end of the file
			break
		}
		if err != nil {
			return "", err
		}
	}

	k := string(content)

	return k, nil
}
