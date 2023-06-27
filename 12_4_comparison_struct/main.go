package main


import "fmt"


type employee struct {

	name   string

	age    int

	salary int

}


func main() {

	emp1 := employee{name: "Sam", age: 31, salary: 2000}

	emp2 := employee{name: "Sam", age: 31, salary: 2000}

	if emp1 == emp2 {

		fmt.Println("emp1 annd emp2 are equal")

	} else {

		fmt.Println("emp1 annd emp2 are not equal")

	}

}
//تایپ های زیر امکان مقایسه را به ما میدهند :
//boolean
//numeric
//string
//pointer
//channel
//interface types
//structs
//array
//
//و اما ۳ تایپ زیر امکان مقایسه را به شما نمیدهند :
//
//Slice
//Map
//Function