package main

import "fmt"

type fake2 struct {
	content2 string
}

func (s *fake2)string2() string {
	return s.content2
}

type fake3 struct {
	content3 string
}

func (s fake3)string3() string {
	return s.content3
}

func printing(val interface{}) {
	//fmt.Println(val.(test).string3())
	result, ok := val.(test)
	if ok{
		fmt.Println(result.string3())
	}
}

type test interface {
	string3() string
}

func main() {
	sample := fake3{content3: "dhfdsfhsdkhfdjh"}
	printing(sample)
}

