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

type Namer interface {
	FullName() string // تو کاری نداشته باش که چه اتفاقی می افتد آبجکتی که قرار است دریافت کنی میبایست متدی به این نام داشته باشد
}

func Greeting(namer Namer) string { // namer : هنگامیکی که از نوع اینترفیس میباشد حتما میبایست آدرس باشد
	return fmt.Sprintf("greeting %s", namer.FullName())
}

func main() {
	//u := User{
	//	Name:   "sajjad",
	//	Family: "moshtaghi",
	//	Age:    25,
	//}
	//fmt.Println(u.FullName())

	user := &User{ // لازم است از نوع آدرس باشد
		Name:   "sajjad",
		Family: "moshtagi",
		Age:    25,
	}
	fmt.Println(Greeting(user))
}
