package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := (a / b); frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	SwitchStatment()
}

func SwitchStatment() {
	x := 2
	y := 5

	// automatically break

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		x = 3
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("many")
	}

	// automatically break
	switch {
	case x > 3:
		fmt.Println("x is greater than 3")
	case y > 3:
		fmt.Println("y is greater than 3")
	case x < 3:
		fmt.Println("x is lester than 3")
	}

	n := 3
	switch n {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
		fallthrough
	case 4:
		fmt.Println("four")
		fallthrough
	case 5:
		fmt.Println("five")
		fallthrough
	case 6:
		fmt.Println("six")
		fallthrough
	default:
		fmt.Println("number")
	}
}
