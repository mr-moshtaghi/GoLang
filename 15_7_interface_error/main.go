package main

import (
	"fmt"
	"time"
)

type test interface { // در کد اصلی گو اینگونه تعریف شده است
	Hello() string
}

type MyError struct {
	CreatedAt time.Time
	Body      string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("At %s Error: %s", me.CreatedAt, me.Body)
}

func (me *MyError) Hello() string {
	return fmt.Sprintf("Atfsdfjjdsfkjdskfjdskfjds %s Error: %s", me.CreatedAt, me.Body)
}

func RunApp() test {
	newError := &MyError{
		CreatedAt: time.Now(),
		Body:      "we have funny error here",
	}
	return newError
}

func main() {
	myApp := RunApp()
	fmt.Println(myApp)
}


//package main
//
//import (
//"fmt"
//"time"
//)
//
//type error interface {  // در کد اصلی گو اینگونه تعریف شده است
//	Error() string
//}
//
//type MyError struct {
//	CreatedAt time.Time
//	Body      string
//}
//
//func (me *MyError) Error() string {
//	return fmt.Sprintf("At %s Error: %s", me.CreatedAt, me.Body)
//}
//
//func RunApp() error {
//	newError := &MyError{
//		CreatedAt: time.Now(),
//		Body:      "we have funny error here",
//	}
//	return newError
//}
//
//func main() {
//	myApp := RunApp()
//	fmt.Println(myApp)
//}

