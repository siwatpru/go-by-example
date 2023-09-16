package main

import "fmt"

type base struct {
  num int
}

func (b base) describe() string {
  return fmt.Sprintf("base with num %v", b.num)
}

// Embedding a struct is like inheriting from a class in OOP.
type container struct {
  base
  str string
}

func main() {
  co := container{
    base: base{num: 1},
    str: "some name",
  }

  // The embedded struct's methods are also promoted to the outer struct.
  fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
  fmt.Println("base: ", co.base.num) // 1
  fmt.Println("describe(): ", co.describe()) // base with num 1
}
