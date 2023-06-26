package main

import "fmt"

type User struct {
	ID     int
	Name   string
	Family string
	Age    int
}

func (u *User) FullName() string {
	return fmt.Sprintf(" %s %s", u.Name, u.Family)
}

func (u *User) GetID() string {
	return fmt.Sprintf("id user is %d", u.ID)
}

type Player struct {
	User
	PlayerID int
}

func (p *Player) FullName() string {
	return fmt.Sprintf("%s %s is %d player", p.Name, p.Family, p.PlayerID)
}

func (p *Player) GetID() string {
	return fmt.Sprintf("id player is %d", p.PlayerID)
}

type Namer interface {
	FullName() string // تو کاری نداشته باش که چه اتفاقی می افتد آبجکتی که قرار است دریافت کنی میبایست متدی به این نام داشته باشد
}

type Identity interface {
	GetID() string
}

type All interface {
	Namer
	Identity
}

func Greeting(namer Namer) string { // namer : هنگامیکی که از نوع اینترفیس میباشد حتما میبایست آدرس باشد
	return fmt.Sprintf("greeting %s", namer.FullName())
}

func ShowID(identity Identity) string {
	return fmt.Sprintf("ID : %s", identity.GetID())
}

func ShowAll(all All) string {
	return fmt.Sprintf("greeting %s, ID is %s", all.FullName(), all.GetID())
}

func main() {

	user := &User{ // لازم است از نوع آدرس باشد
		Name:   "sajjad",
		Family: "moshtagi",
		Age:    25,
	}

	player := &Player{
		User: User{
			Name:   "ahmad",
			Family: "safa",
			Age:    25,
		},
		PlayerID: 5,
	}
	fmt.Println(Greeting(user))
	fmt.Println(Greeting(player))

	fmt.Println(ShowID(user))
	fmt.Println(ShowID(player))

	fmt.Println(ShowAll(user))
	fmt.Println(ShowAll(player))
}
