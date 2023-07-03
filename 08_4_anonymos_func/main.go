// Go program to illustrate how
// to create an anonymous function
//package main
//
//import "fmt"
//
//func main() {
//
//	// Anonymous function
//	func(){
//
//		fmt.Println("my name is sajjad")
//	}()
//}

// ##################################################################################

// Go program to illustrate
// use of an anonymous function
//package main
//
//import "fmt"
//
//func main() {
//
//	// Assigning an anonymous
//	// function to a variable
//	value := func(){
//		fmt.Println("my name is sajjad")
//	}
//	value()
//
//}

// ######################################################################################

//// Go program to pass arguments
//// in the anonymous function
//package main
//
//import "fmt"
//
//func main() {
//
//	// Passing arguments in anonymous function
//	func(ele string) {
//		fmt.Println(ele)
//	}("my name is sajjad")
//
//}

// ######################################################################################

// Go program to pass an anonymous
// function as an argument into
// other function
package main

import "fmt"

// Passing anonymous function
// as an argument
func GFG(i func(p, q string) string) {
	fmt.Println(i("my ", "name is "))

}

func main() {
	value := func(p, q string) string {
		return p + q + "sajjad"
	}
	GFG(value)
}

// ######################################################################################
