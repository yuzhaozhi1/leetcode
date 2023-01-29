package main

var (
	result [][]int
	path   []int
)

// 只使用数字1到9
func combinationSum3(k int, n int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, k)
	dfs(k, n, 1, 0)
	return result
}

func dfs(k, n, start, sum int) {
	if len(path) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
		}
		return
	}

	for i := start; i <= 9 && i < n-(k-len(path))+1; i++ {
		path = append(path, i)
		dfs(k, n, i+1, sum+i)
		path = path[:len(path)-1]
	}
}
