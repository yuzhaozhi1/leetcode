package main

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/

var (
	// 电话号码映射      0    1    2      3     4      5       6      7       8       9
	tmp = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "xyz"}
)

func letterCombinations(digits string) []string {
	res := make([]string, 0, len(digits))
	if digits == "" {
		return res
	}

}
