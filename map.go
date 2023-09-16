package main

import (
  "fmt"
  // "maps"
)

func main() {

  m := make(map[string]int)
  fmt.Println("map:", m) // map: map[]

  // Set key/value pairs
  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("map:", m) // map: map[k1:7 k2:13]

  v1 := m["k1"]
  fmt.Println("v1:", v1) // v1: 7
  v3 := m["k3"]
  fmt.Println("v3:", v3) // v3: 0

  fmt.Println("len:", len(m)) // len: 2

  // Delete key/value pairs
  delete(m, "k2")
  fmt.Println(m) // map[k1:7]

  k1, prs1 := m["k1"]
  fmt.Println("k1:", k1) // k1: 7
  fmt.Println("prs1:", prs1) // prs1: true

  k2, prs2 := m["k2"]
  fmt.Println("k2:", k2) // k2: 0
  fmt.Println("prs2:", prs2) // prs2: false

  // Clear the map
  clear(m)
  fmt.Println(m) // map[k1:7]

  // Shorthand declaration
  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n) // map: map[foo:1 bar:2]
}
