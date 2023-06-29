package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	nums := []int{2, 2, 1, 4, 4, 1, 6}
	result := singleNumber(nums)
	fmt.Println(result)
}
