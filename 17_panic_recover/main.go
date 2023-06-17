package main

import "fmt"

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR : %s\n", err)
		}
	}()

	return vals[index]
}

func main() {
	// val := []int{1, 2, 3}
	// // panic
	// v := val[5]
	// fmt.Println(v)

	// file, err := os.Open("no-directory")
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()
	// fmt.Println("file oppened")

	v := safeValue([]int{1, 2, 3}, 5)
	fmt.Println(v)
}