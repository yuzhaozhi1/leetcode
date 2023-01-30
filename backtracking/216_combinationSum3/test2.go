package main

var (
	res   [][]int
	tmper []int
)

// 从1~9 取数 k为数组中的数, n 为sum 和
func combinationSum4(k int, n int) [][]int {
	res, tmper = make([][]int, 0), make([]int, 0, k)
	backtracking(k, n, 1, 0)
	return res
}

func backtracking(k, n, startIndex, sum int) {
	if len(tmper) == k {
		if sum == n {
			tmpVar := make([]int, k)
			copy(tmper, tmpVar)
			res = append(res, tmpVar)
		}
		return
	}
	for i := startIndex; i <= 9 && i < 9-(k-len(tmper))+1; i++ {
		sum += i
		path = append(path, i)
		backtracking(k, n, i, sum)
		path = path[:len(path)-1]
	}
}
