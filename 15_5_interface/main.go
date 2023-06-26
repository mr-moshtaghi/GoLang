package main

import "fmt"

type User struct {
	Name   string
	Family string
	Age    int
}

func (u *User) FullName() string {
	return fmt.Sprintf(" %s %s", u.Name, u.Family)
}

type Player struct {
	User
	PlayerID int
}

func (p *Player) FullName() string {
	return fmt.Sprintf("%s %s is %d player", p.Name, p.Family, p.PlayerID)
}

type Namer interface {
	FullName() string // تو کاری نداشته باش که چه اتفاقی می افتد آبجکتی که قرار است دریافت کنی میبایست متدی به این نام داشته باشد
}

func Greeting(namer Namer) string { // namer : هنگامیکی که از نوع اینترفیس میباشد حتما میبایست آدرس باشد
	return fmt.Sprintf("greeting %s", namer.FullName())
}

func main() {

	user := &User{ // لازم است از نوع آدرس باشد
		Name:   "sajjad",
		Family: "moshtagi",
		Age:    25,
	}

	player := &Player{
		User:     User{
			Name:   "ahmad",
			Family: "safa",
			Age:    25,
		},
		PlayerID: 5,
	}
	fmt.Println(Greeting(user))
	fmt.Println(Greeting(player))
}
