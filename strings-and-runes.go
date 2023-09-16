package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "こんにちは, 世界"
	fmt.Println(s) // こんにちは, 世界
  fmt.Println("rune count:", utf8.RuneCountInString(s)) // rune count: 9

  // range iterates over string and returns runes and indexes
  for idx, runeValue := range s {
    fmt.Printf("%#U starts at byte position %d\n", runeValue, idx)
  }

  // DecodeRuneInString decodes the first UTF-8 encoding in s and returns the rune and its width in bytes
  fmt.Println("\nUsing DecodeRuneInString")
  for i, w := 0, 0; i < len(s); i += w {
    runeValue, width := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
    w = width
    if runeValue == '世' {
      fmt.Println("found 世")
    }
  }
}
