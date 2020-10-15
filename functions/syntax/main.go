package main

import "fmt"

func main() {
	foo()
	bar("james")
	s1 := woo("MOneyPenney")
	fmt.Println(s1)
	x, y := mouse("IAn", "fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("hey there")
}

// everthing in go is pass by value
func bar(s string) {
	fmt.Println("Hi therer ", s)
}

func woo(st string) string {
	return fmt.Sprintln("hello from woo", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprintln(fn, "", ln, `, says "Hello"`)
	b := true
	return a, b
}
