package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("reader_interface/demo_txt.txt")

	if err != nil {
		log.Fatal("couldn't open file")
	}

	l := 10

	buf := make([]byte, l)

	fmt.Printf("Contents of buf before reading file: %v\n", buf)

	n, err := f.Read(buf)

	fmt.Printf("Contents of buf after reading file: %v\n", buf)

	j := buf[:n]

	fmt.Printf("Contents of buf through n bytes: %v\n", j)

	k := string(j)

	fmt.Printf("Contents of buf through n bytes as string: %v\n", k)

}
