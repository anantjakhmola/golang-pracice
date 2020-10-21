package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{3, 1, 2, 4, 1, 31, 2}
	sort.Ints(s)
	fmt.Println(s)

	s1 := []string{"James", "Q", "M", "Moneypenny"}
	sort.Strings(s1)
	fmt.Println(s1)
}
