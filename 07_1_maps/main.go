package main

import "fmt"

func main() {
	bar := map[string]float64{
		"Amazon":    1700.0,
		"Google":    1130.0,
		"Microsoft": 100.5,
	}

	fmt.Println(bar)

	fmt.Println("====================================")

	// get value
	fmt.Println(bar["Amazon"])
	fmt.Println("====================================")

	//get zero value if not fount
	fmt.Println(bar["Apple"])
	fmt.Println("====================================")

	//use to value to see if found
	value, ok := bar["Apple"]

	if !ok {
		fmt.Println("Apple not found")
	} else {
		fmt.Println(value)
	}
	fmt.Println("====================================")

	//set
	bar["Apple"] = 350.5
	fmt.Println(bar)
	fmt.Println("====================================")

	//delete
	delete(bar, "Amazon")
	fmt.Println(bar)
	fmt.Println("====================================")

	// singel value "for" is on keys
	for key := range bar {
		fmt.Println(key)
	}
	fmt.Println("===================================")

	// singel value "for" is key and value
	for key, value := range bar {
		fmt.Printf("%s - %.2f\n", key, value)
	}
	fmt.Println("===================================")

}
