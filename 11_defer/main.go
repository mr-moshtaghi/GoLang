package main

import "fmt"

func cleanup(name string) {
	fmt.Printf("cleaning up %s\n", name)
}

func worker() {
	fmt.Println("worker")
}

func main() {
	defer cleanup("A")
	defer cleanup("B")

	worker()
}
