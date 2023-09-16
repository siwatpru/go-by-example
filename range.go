package main

import "fmt"

func main() {
  nums := []int{2, 3, 4}
  sum := 0

  // range returns both the index and value
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum) // sum: 9

  // print index of 3 in the slice
  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i) // index: 1
    }
  }

  // iterate over key/value pairs in a map
  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

  // iterate over keys in a map
  for k := range kvs {
    fmt.Println("key:", k)
  }

  // iterate over unicode code points in a string
  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
