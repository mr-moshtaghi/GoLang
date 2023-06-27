package main

import "fmt"

func main() {
	b := [][]string{{"this", "is", "first"}, {"this", "is"}}
	b[1][1] = "second"
	fmt.Println(b)

	c := [][]string{{"this", "is", "first"}, {"this", "is"}}
	fmt.Println(c)
	c[0][2] = "second"
	fmt.Println(c)

	c[1] = append(c[1], "second")
	fmt.Println(c)

	c = append(c, []string{"third"})
	fmt.Println(c)

	citiesLen := 10
	cities := make([]int, citiesLen)
	fmt.Println(cities[:5])

	p := []int{}
	p = append(p, 10, 20, 30)
	fmt.Println(p)
}
