package main

import "fmt"

type fake2 struct {
	content2 string
}

func (s fake2)string2() string {
	return s.content2
}

type fake3 struct {
	content3 string
}

func (s fake3)string3() string {
	return s.content3
}

func printing(val interface{}) {
	switch rt := val.(type) {
	case test2:
		fmt.Println(rt.string2())
	case test3:
		fmt.Println(rt.string3())
	}
}

type test3 interface {
	string3() string
}

type test2 interface {
	string2() string
}

func main() {
	sample3 := fake3{content3: "test for fake 3"}
	sample2 := fake2{content2: "test for fake 2"}
	fmt.Println(sample3.string3())
	printing(sample2)
}

