package main

import "os"

func main() {

  // Panic is a built-in function that stops
  // the ordinary flow of control and
	panic("a problem")

  // A common use of panic is to abort if a function
  // returns an error value that we don’t know how to
	_, err := os.ReadFile("/tmp/file-xxx-x")
	if err != nil {
		panic(err)
	}
}
