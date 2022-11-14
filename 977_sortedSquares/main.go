package main

func sortedSquares(nums []int) []int {

	tmp := make([]int, len(nums))
	i := 0
	j := len(nums) - 1
	k := len(nums) - 1
	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			tmp[k] = nums[i] * nums[i]
			i++
		} else {
			tmp[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return tmp
}
