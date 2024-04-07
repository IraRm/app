package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scanln(&n)
	for i := 1.0; i <= n; i++ {
		fmt.Println(math.Pow(2, i))
	}
}
