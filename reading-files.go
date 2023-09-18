package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Note: Assuming file in /tmp/dat exists
// echo 'hello' > /tmp/dat
// echo 'go' >>   /tmp/dat
// go run reading-files.go

func main() {
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat")
	check(err)

	// Read some bytes from the beginning of the file.
	b1 := make([]byte, 5) // 5 bytes
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Seek to a known location in the file and Read from there.
	o2, err := f.Seek(6, 0) // From 6th bytes
	check(err)
	b2 := make([]byte, 2) // 2 bytes
	n2, err := f.Read(b2) // Read 2 bytes from 6th bytes
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Rewind to the beginning of the file
	_, err = f.Seek(0, 0)
	check(err)

	// The bufio package implements a buffered reader
	// that may be useful both for its efficiency with
	// many small reads and because of the additional
	// reading methods it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
