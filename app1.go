package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 4, 5}
	numsRev := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		numsRev = append(numsRev, (nums[i]))
	}
	fmt.Println(numsRev)
}
