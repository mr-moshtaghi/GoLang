package main

import "fmt"

// oop

type Trade struct {
	// if property  starts with capital letter, that is public otherwise is private

	Symbol string  // stock symbol
	Volume int     // number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade

}

func main() {
	t1 := Trade{"Apple", 10, 99.98, true}

	fmt.Println(t1)

	fmt.Println(t1.Symbol)

	fmt.Printf("%+v", t1)

	fmt.Println("========================================================")

	t2 := Trade{
		Price:  99.97,
		Symbol: "Microsoft",
		Volume: 15,
		Buy:    false, // trailing comma is mandatory
	}

	println(t2)
	fmt.Printf("%+v", t2)

	fmt.Println("========================================================")

	t3 := Trade{}
	fmt.Println(t3)
	fmt.Printf("%+v", t3)

}
