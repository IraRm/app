package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	s = ReverseLetters(s)
	fmt.Println(s)
}
func ReverseLetters(s string) (res string) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {
			res += string(s[i])
		}
	}
	return
}
