package main

import (
	"encoding/json"
	"fmt"
)

//! Here I think that by addding json  
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"BOnd","Age":21},{"First":"Mary","Last":"mob","Age":30}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all data", people)
	for i, v := range people {
		fmt.Println("\nPERSON", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
