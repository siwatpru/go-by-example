package main

import (
  "fmt"
  "time"
)

func f(from string) {
  for i := 0; i < 3; i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {

  // call the function directly
  f("direct")

  // call the function in a goroutine
  go f("goroutune1")
  go f("goroutune2")

  // call the anonymous function in a goroutine
  go func(msg string) {
    fmt.Println(msg)
  }("going")

  // wait for the goroutines to finish
  // PS: Should use WaitGroup instead
  time.Sleep(time.Second)
  fmt.Println("done")
}
