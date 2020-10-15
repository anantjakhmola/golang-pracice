package main

import "fmt"

type vehicle struct {
	door  string
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	c := truck{
		vehicle: vehicle{
			door:  "4",
			color: "blue",
		},
		fourwheel: true,
	}
	c1 := sedan{
		vehicle: vehicle{
			door:  "4",
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(c, c1)

}
