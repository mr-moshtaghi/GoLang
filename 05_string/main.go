package main

import "fmt"

func main() {
	book := "my name is sajjad"

	fmt.Println(book)

	// string length
	fmt.Println(len(book))

	// slice
	fmt.Println(book[4:11])
	fmt.Println(book[4:])
	fmt.Println(book[:4])

	// use + to concatinate string
	fmt.Println("hey " + book)

	//multi line
	poem := `
	hello 
	how are you?
	`

	fmt.Println(poem)

	fmt.Printf("book[0] =  %v and type(book[0]) = %T", book[0], book[0])
}
