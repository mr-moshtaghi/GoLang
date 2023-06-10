package main

import (
	"fmt"
	"reflect"
)

// var name string
// var age int
// var location string

// var(
// 	name string
// 	age int
// 	location string
// )

// var (
// 	name, location string
// 	age            int
// )

// var (
// 	name, location string
// 	age            int = 12
// )

var (
	name, location string
	age            = 12
)

func main() {
	family := 12
	// print(age)
	fmt.Println(reflect.TypeOf(family))
}
