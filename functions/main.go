package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divideModular(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	val_add := add(3, 4)
	fmt.Println(val_add)

	div, mod := divideModular(7, 2)
	fmt.Printf("div = %d - mod= %d", div, mod)
	// fmt.Println(div, mod)
}
