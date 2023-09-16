package main

import "fmt"

// MapKeys returns a slice of keys from a map
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// List is a singly-linked list
type List[T any] struct {
	head, tail *element[T]
}

// element is a node in a singly-linked list
type element[T any] struct {
  next *element[T]
  val T
}

// Push adds a value to the end of the list
func (lst *List[T]) Push(v T) {
  if lst.tail == nil {
    lst.head = &element[T]{val: v}
    lst.tail = lst.head
  } else {
    lst.tail.next = &element[T]{val: v}
    lst.tail = lst.tail.next
  }
}

// GetAll returns all elements in the list as a slice
func (lst *List[T]) GetAll() []T {
  var elements []T
  for e := lst.head; e != nil; e = e.next {
    elements = append(elements, e.val)
  }
  return elements
}

func main() {
  var m = map[int]string{1: "2", 2: "4", 4: "8"}
  fmt.Println("keys: ", MapKeys(m))

  n := MapKeys[int, string](m)
  fmt.Println(n)

  lst := List[int]{}
  lst.Push(1)
  lst.Push(2)
  lst.Push(3)
  fmt.Println("list: ", lst.GetAll())
}
