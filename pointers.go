package main

import "fmt"

func zeroval(i int) {
  i = 0
}

func zeroptr(i *int) {
  *i = 0
}

func main() {
  i := 1
  fmt.Println("initial:", i) // 1

  zeroval(i)
  fmt.Println("zeroval:", i) // 1

  zeroptr(&i)
  fmt.Println("zeroptr:", i) // 0

  fmt.Println("pointer:", &i) // pointer value
}
