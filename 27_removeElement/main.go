package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3, 4}
	element := removeElement(nums, 3)
	fmt.Println(element)
}

func removeElement(nums []int, val int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		// 如果 快指针的值不等于 val
		// 就将数据放到慢指针所指的位置
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}
	// nums = nums[:slow]
	return slow
}
