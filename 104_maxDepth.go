package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepth(root *TreeNode) int {
	res := 0
	var order func(root *TreeNode, level int)
	order = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		res = max(res, level)
		if root.Left != nil {
			order(root.Left, level+1)
		}
		if root.Right != nil {
			order(root.Right, level+1)
		}
	}
	order(root, 1)
	return res
}

func main1() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	depth := maxDepth(root)
	fmt.Println(depth)
}
