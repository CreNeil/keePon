package main

import "sort"

func removeDuplicates(nums []int) int {
	sort.Ints(nums)
	location := 0
	if len(nums) != 0 {
		for i := 0; i < len(nums); i++ {
			if nums[i] > nums[location] {
				nums[location+1] = nums[i]
				location++
			}
		}
	}
	return location + 1
}

func main() {
	nums := []int{4, 2, 3, 2, 1, 1, 1, 2}
	println(removeDuplicates(nums))
}
