package main

import (
	"fmt"
	"time"
)

type NewStruct struct {
	Name        string `json:"name"`
	Family      string
	HomeAddress string `json:"home_address"`
	Model       NewModel
}

type NewModel struct {
	ID        uint64
	CreatedAt time.Time
}

func (ns *NewModel) SayHello() string {
	return "hello"
}

func main() {
	new_struct := NewStruct{
		Name:        "sajjad",
		Family:      "moshtaghi",
		HomeAddress: "ardestan",
	}
	fmt.Println(new_struct.Name)
	fmt.Println(new_struct.Model.ID)
}
