
package main

import "fmt"

type User struct {
	Name   string
	Family string
	age    int64
}

func (u User) PrintingName() {
	fmt.Println(u.Name)
}

type Player struct {
	User
	PlayerId int64
}

func main() {
	p := Player{
		User: User{
			Name:   "sajjad",
			Family: "moshtaghi",
			age:    25,
		},

		PlayerId: 20,
	}
	fmt.Println(p.Family)
	p.PrintingName()
}
