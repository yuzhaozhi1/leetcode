package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// preorderTraversal 前序遍历  144
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 5)
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return res
}

// inorderTraversal 中序遍历 94
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 5)
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)
	}
	traversal(root)
	return res
}

// postorderTraversal 145 后序遍历
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 5)
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		traversal(root.Right)
		res = append(res, root.Val)
	}
	traversal(root)
	return res
}
