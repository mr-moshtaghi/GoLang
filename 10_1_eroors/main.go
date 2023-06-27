package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 { // if negetive
		return 0.0, fmt.Errorf("sqrt of negetive value (%f)", n)
	}
	return math.Sqrt(n), nil
}

func main() {
	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("error: %s \n", err)
	} else {
		fmt.Println(s1)
	}

	s2, err := sqrt(-2.0)
	if err != nil {
		fmt.Printf("error: %s \n", err)
	} else {
		fmt.Println(s2)
	}
}
