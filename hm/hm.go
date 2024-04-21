package main

import "fmt"

func main() {
	var x, y int
	var opr string
	fmt.Scan(&x, &y, &opr)

	if opr != "*" || opr != "+" || opr != "-" || opr != "/" {
		fmt.Println("Неверная операция")
	}
	if opr == "+" {
		fmt.Println(x + y)
	} else if opr == "-" {
		fmt.Println(x - y)
	} else if opr == "*" {
		fmt.Println(x * y)
	} else if opr == "/" {
		if y != 0 {
			fmt.Println(x / y)
		} else {
			fmt.Println("На ноль делить нельзя!")
		}
	}
}
