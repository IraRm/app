package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums = change(nums)
	fmt.Println(nums)
}
func change(nums []int) []int {
	for i := 0; i <= len(nums); i++ {
		nums[i] = nums[len(nums)-1]
	}
	return nums
}
