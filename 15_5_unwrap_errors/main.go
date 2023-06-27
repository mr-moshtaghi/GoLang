package main

import (
	"errors"

	"fmt"
)

type errorOne struct{}

func (e errorOne) Error() string {

	return "Error One happened"

}

func main() {

	e1 := errorOne{}

	e2 := fmt.Errorf("E2: %w", e1)

	e3 := fmt.Errorf("E3: %w", e2)

	fmt.Println(errors.Unwrap(e3))

	fmt.Println(errors.Unwrap(e2))

	fmt.Println(errors.Unwrap(e1))

}


// در کد بالا متغیر e2 خطای داخل ساختار e1 را هم پوشانی کرده و سپس متغیر e3 خطای متغیر e2
//را هم پوشانی می کند. در نهایت با تابع Unwrap متن خطای اصلی را چاپ کردیم.