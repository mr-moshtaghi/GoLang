package main

import (
	"fmt"
)

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR : %s\n", err)
		}
	}()

	return vals[index]
}

func main() {
	//val := []int{1, 2, 3}
	//panic("hello")
	//v := val[5]
	//fmt.Println(v)

	//file, err := os.Open("no-directory")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer file.Close()
	//fmt.Println("file oppened")

	v := safeValue([]int{1, 2, 3}, 5)
	fmt.Println(v)
}




//package main
//
//
//import "fmt"
//
//
//func main() {
//
//
//	a := []string{"a", "b"}
//
//	checkAndPrint(a, 2)
//
//}
//
//
//func checkAndPrint(a []string, index int) {
//
//	if index > (len(a) - 1) {
//
//		panic("Out of bound access for slice")
//
//	}
//
//	fmt.Println(a[index])
//
//}
