package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolates"}
	fmt.Println(jb)
	mp := []string{"Miss", "MOneypenny", "struts"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
