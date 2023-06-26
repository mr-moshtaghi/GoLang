package main

import "fmt"

type T struct {
	Name string
}

type User struct {
	FirstName, LastName string
}

func (u User)Greeting() string {
	u.FirstName = "ahmad"
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

//func (u *User)Greeting() string {
//	u.FirstName = "ahmad"
//	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
//}

func main() {
	p := &T{
		Name: "quotMohammad",
	}
	fmt.Printf("p is pointing to: %p \n", p) // آدرسی که p به آن اشاره میکند
	fmt.Printf("p address: %p \n", &p)  // آدرس خود p

	newP := p
	fmt.Printf("newP is pointing to: %p \n", newP)
	fmt.Printf("newP address: %p \n", &newP)


	u := User{
		FirstName: "sajjad",
		LastName:  "moshtaghi",
	}
	fmt.Println(u.Greeting())
	fmt.Println(u)

}
