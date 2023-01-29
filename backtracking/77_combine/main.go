package main

var (
	result [][]int
	path   []int
)

// combine 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// n 是从1到n 在数组中可以取到的最大的数
// k 为数组需要取几个数
func combine(n int, k int) [][]int {
	result = make([][]int, 0, n)
	path = make([]int, 0, k)

	backtrackingCombine(n, k, 1) // 1 是因为,题目给的就是 范围为 [1,n]
	return result
}

func backtrackingCombine(n, k, startIndex int) {
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		result = append(result, tmp)
		return
	}
	// len(path) 已经选择的数
	// k-len(path) 还需要添加的数
	// 在集合n中至多要从该起始位置 : n - (k - path.size()) + 1，开始遍历
	// 为什么有个+1呢，因为包括起始位置，我们要是一个左闭的集合。
	// i为本次搜索的起始位置
	for i := startIndex; i <= n-(k-len(path))+1; i++ { // 从start开始，不往回走，避免出现重复组合
		path = append(path, i)
		backtrackingCombine(n, k, i+1)
		path = path[:len(path)-1]
	}
}

func dfs(n int, k int, start int) {
	if len(path) == k { // 说明已经满足了k个数的要求
		tmp := make([]int, k)
		copy(tmp, path)
		result = append(result, tmp)
		return
	}
	for i := start; i <= n; i++ { // 从start开始，不往回走，避免出现重复组合
		if n-i+1 < k-len(path) { // 剪枝
			break
		}
		path = append(path, i)
		dfs(n, k, i+1)
		path = path[:len(path)-1]
	}
}

//
func combine2(n int, k int) [][]int {
	result = make([][]int, 0, n)
	path = make([]int, 0, k)
	backtrackingCombine2(n, k, 1)
	return result
}

// startIndex 表示从那个数开始取,  比如n 为4, start 为2, 则表示可以取 2,3,4 这三个数
func backtrackingCombine2(n, k, startIndex int) {
	// 结束判断
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := startIndex; i < n-(k-len(path))+1; i++ {
		// 继续遍历
		backtrackingCombine2(n, k, startIndex+1)
		// 每次遍历后都需要将之前path 中的数据给丢弃
		path = path[:len(path)-1]
	}
}

func main() {
	combine(4, 2)
}
