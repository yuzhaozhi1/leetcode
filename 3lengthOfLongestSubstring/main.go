package main

import (
	"fmt"
	"strings"
)

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

func lengthOfLongestSubstring(s string) int {
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i])) // 找到当前的i
		if index == -1 && i+1 > end {
			end = i + 1
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

func main() {
	s := "abcabcbb"
	substring := lengthOfLongestSubstring(s)
	fmt.Println(substring)
}
