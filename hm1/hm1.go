package main

import "fmt"

func main() {
	for i := 10; i < 100; i++ {
		if ((i/10)*(i%10))*2 == i {
			fmt.Println(i)
		}
	}
}
