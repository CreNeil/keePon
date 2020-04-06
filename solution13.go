package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	begin := make([]int, len(nums))
	begin[0] = nums[0]
	result := begin[0]
	for i := 1; i < len(nums); i++ {
		begin[i] = max(nums[i], nums[i]+begin[i-1])
		result = max(result, begin[i])
	}
	return result
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	print(maxSubArray(nums))
}
