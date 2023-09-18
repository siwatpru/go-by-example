package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type repsonse1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// JSON encoding exmaples
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	// JSON encoding from struct
	res1D := &repsonse1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

  // JSON encoding from struct with specified field names
  res2D := &response2{
    Page:   1,
    Fruits: []string{"apple", "peach", "pear"},
  }
  res2B, _ := json.Marshal(res2D)
  fmt.Println(string(res2B))

  // JSON decoding to Go values
  byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

  // We need to provide a variable where the JSON
  // package can put the decoded data. This map[string]interface{}
  // will hold a map of strings to arbitrary data types.
  var dat map[string]interface{}

  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  fmt.Println(dat)

  // Accessing nested data requires a series of casts.
  num := dat["num"].(float64)
  fmt.Println(num)
  strs := dat["strs"].([]interface{})
  str1 := strs[0].(string)
  fmt.Println(str1)

  // We can also decode JSON into custom data types. This has
  // the advantages of adding additional type-safety to our
  // programs and eliminating the need for type assertions
  // when accessing the decoded data.
  str := `{"page": 1, "fruits": ["apple", "peach"]}`
  res := response2{}
  json.Unmarshal([]byte(str), &res)
  fmt.Println(res)
  fmt.Println(res.Fruits[0])

  // Stream JSON encodings directly to os.Writers like
  // os.Stdout or even HTTP response bodies.
  enc := json.NewEncoder(os.Stdout)
  d := map[string]int{"apple": 5, "lettuce": 7}
  enc.Encode(d)
}
