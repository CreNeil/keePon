package main

import "sort"

func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	begin := -1
	end := -1
	if len(nums) != 0 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == val && begin == -1 {
				begin = i
			}
			if nums[i] == val && begin != -1 {
				end = i
			}
		}
	}
	if begin != -1 && end != -1 {
		nums = append(nums[:begin], nums[end+1:]...)
	}
	return len(nums)
}

func main() {
	nums := []int{4, 2, 3, 2, 1, 1, 1, 2, 4}
	println(removeElement(nums, 4))
}
