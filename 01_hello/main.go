package main

import (
	// strutils "GoLang/pakages"
	"GoLang/packages"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	// fmt.Println(strutils.Revers("hello world"))
	fmt.Println(pakages.Revers("hello world"))
	fmt.Println(pakages.Pi)
	value, _ := pakages.SayHello("sajjad", "moshtaghi")
	fmt.Println(value)
}
