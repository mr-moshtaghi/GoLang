package main

import "fmt"

func doubleAt(values []int, i int) { // slice , map pass by reference
	values[i] *= 2
}

func double(n int) { // passed by value (int)
	n *= 2
}

func doublePtr(n *int) { // pointer
	*n *= 2
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	doubleAt(values, 2)
	fmt.Println(values)
	fmt.Println("========================================")

	val := 10
	double(val)
	fmt.Println(val)
	fmt.Println("========================================")

	doublePtr(&val) //pointer
	fmt.Println(val)
	fmt.Println("========================================")

}
