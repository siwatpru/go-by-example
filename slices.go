package main

import (
	"fmt"
	"slices"
)

func main() {

  // Create a string array
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0) // len: 0, cap: 0

  // Make a string slices with size 3
	s = make([]string, 3) 
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // len: 3, cap: 3 

  // Set value of the slices
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s) // [a b c]
	fmt.Println("get:", s[2]) // c
	fmt.Println("len:", len(s)) // 3

  // Append value to the slices
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s, len(s), cap(s)) // [a b c d e f] 6 6


  // Create a new slice with the same size as s
	c := make([]string, len(s))
	copy(c, s) // copy s to c
	fmt.Println("cpy:", c) // [a b c d e f]

  // Get a slice of the slice
	l := s[2:5]
	fmt.Println("sl1:", l) // [c d e]

	l = s[:5]
	fmt.Println("sl2:", l) // [a b c d e]

	l = s[2:]
	fmt.Println("sl3:", l) // [c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) // [g h i]

  // Compare two slices
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

  // Create a 2d slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
