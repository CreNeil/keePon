package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dif := target - nums[i]
		location, ok := result[dif]
		if ok {
			return []int{location, i}
		}
		result[nums[i]] = i
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{3, 2, 4}
	nums = twoSum(nums, 6)
	fmt.Println(nums)
}
