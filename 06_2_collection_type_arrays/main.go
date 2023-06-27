package main

import "fmt"

func main() {
	a := [2][3]string{{"this", "is", "first"}, {"this", "is"}}
	a[1][2] = "second"
	fmt.Println(a)
}
