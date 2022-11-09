package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3, 4, 3, 4, 3}
	element := removeElement(nums, 3)
	fmt.Println(element)
}

func removeElement(nums []int, val int) int {
	i := 0
	length := len(nums)
	j := length - 1
	for i <= j {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
			i++
		} else {
			i++
		}
	}
	fmt.Println(nums)
	return j
}
