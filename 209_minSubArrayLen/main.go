package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	i := 0                           // 滑动窗口开始的位置
	result := len(nums) + 1          // 返回的最大的值
	sum := 0                         // 滑动窗口的总和
	subLength := 0                   // 滑动窗口的长度
	for j := 0; j < len(nums); j++ { // j为滑动窗口结束的位置
		sum += nums[j]
		for sum >= target { // 如果滑动窗口内的总和要大于目标值, 这个时候我们就需要移动滑动窗口开始的下标位置
			subLength = j - i + 1
			if subLength < result {
				result = subLength
			}
			sum = sum - nums[i]
			i++
		}
	}
	if result == len(nums)+1 {
		return 0
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	arrayLen := minSubArrayLen(15, nums)
	fmt.Println(arrayLen)
}
