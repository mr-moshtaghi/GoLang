package main

import (
	"fmt"
	"strings"
)

type MyString string

func (s MyString) Uppercase() string {
	return strings.ToUpper(string(s))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	my_string := MyString("sallam")
	fmt.Println(my_string.Uppercase())

	my_float := MyFloat(-65)
	fmt.Println(my_float.Abs())
}

// alias : چیزی که وجود دارد و ما به آن نامی مستعار میدهیم و بر اساس نیاز حود آن را تغییر میدهیم
