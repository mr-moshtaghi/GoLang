package main

import (
	"fmt"
)

func main() {
	names := []string{"ahmad", "ali", "sajjad", "hasan"}

	// length
	fmt.Println(len(names))
	fmt.Println("===================================")

	//index
	fmt.Println(names[1])
	fmt.Println("===================================")

	//slice
	fmt.Println(names[:2])
	fmt.Println("===================================")

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println("===================================")

	// single value range(to show only index)
	for i := range names {
		fmt.Println(i)
	}
	fmt.Println("===================================")

	// double value range (to show index and value)
	for i, name := range names {
		fmt.Println(i, name)
	}
	fmt.Println("===================================")

	// double value range (to show index and value)
	for i, name := range names {
		fmt.Printf("%d - %s\n", i, name)
	}
	fmt.Println("===================================")

	// double value range (ignore index)
	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Println("===================================")

	// Append
	names = append(names, "mahmud")
	test := append(names, "zahra")
	fmt.Println(names)
	fmt.Println(test)
	fmt.Println("===================================")

}
