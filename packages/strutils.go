// package strutils
package pakages

import "fmt"

var Pi = 3.14

func Revers(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SayHello(name string, family string) (result string, err error)  {
	result = fmt.Sprintf("Hello %s %s", name, family)
	return
}