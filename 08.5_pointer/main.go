package main

import "fmt"

var family = "hassani"

func main() {
	family2 := family
	fmt.Println(family2)

	family2 = "moshtaghi"
	fmt.Println(family2)

	fmt.Println(family)

	fmt.Println("===================================================")

	family3 := &family
	fmt.Println(*family3)

	*family3 = "khorsand"

	fmt.Println(*family3)
	fmt.Println(family)


}
