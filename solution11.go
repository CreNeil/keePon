package main

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target || nums[i] > target {
			return i
		}
	}

	return len(nums)
}
func searchInsertMid(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
func main() {
	nums := []int{1, 2, 3, 4, 6, 7}
	print(searchInsertMid(nums, 8))
}
