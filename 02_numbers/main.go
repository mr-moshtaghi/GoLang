package main

import "fmt"

func main() {
	// var x int
	// var y int

	// x := 2.0
	// y := 1.0

	x := 2
	y := 1

	const test = "sajjad"


	fmt.Printf("x = %v, type of %T\n", x, x)
	fmt.Printf("y = %v, type of %T\n", y, y)

	var mean float64

	// mean = (x + y) / 2.0

	mean = float64(x+y) / 2.0

	fmt.Printf("mean = %v, type of %T\n", mean, mean)
	fmt.Printf("test = %v, type of %T\n", test, test)

}
