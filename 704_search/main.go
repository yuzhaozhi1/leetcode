package main

func search(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1
	for left <= right {
		mid := ((right - left) >> 1) + left
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	length := len(nums)
	i := 0
	j := length - 1
	for i <= j {
		if nums[i] == target {
			return i
		}
		if nums[j] == target {
			return j
		}
		j--
		i++
	}
	return -1
}
