package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

// constructor
func newTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol can not be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("price must be >= 0 (was %f)", price)
	}

	trade := Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return &trade, nil
}

func (trade *Trade) Value() float64 {
	value := float64(trade.Volume) * trade.Price
	if trade.Buy {
		value = -value
	}
	return value
}

func main() {
	t, err := newTrade("Microsoft", 10, -100.00, true)
	if err != nil {
		fmt.Printf("ERROR: can not create trade - %s\n", err)
		os.Exit(1)
	}
	fmt.Println(t.Value())

}
