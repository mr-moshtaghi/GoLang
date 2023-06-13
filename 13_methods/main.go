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

func (trade *Trade) Value() float64 {
	value := float64(trade.Volume) * trade.Price
	if trade.Buy {
		value = -value
	}
	return value
}

type Point struct {
	x int
	y int
}

func (point Point) Move(dx int, dy int) {
	point.x += dx
	point.y += dy

}

func main() {
	t := Trade{"Apple", 10, 99.98, true}

	fmt.Println(t.Value())

	point := Point{
		x: 1,
		y: 2,
	}
	point.Move(2, 3)
	fmt.Printf("%+v\n", point)

}
