package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var res [][]string
	var order func(root *TreeNode, tmp *[]string, all *[][]string)

	order = func(root *TreeNode, tmp *[]string, res *[][]string) {
		*tmp = append(*tmp, strconv.Itoa(root.Val))

		if root.Left == nil && root.Right == nil {
			*res = append(*res, *tmp)
		}

		if root.Left != nil {
			newTmp := make([]string, 0, len(*tmp))
			newTmp = append(newTmp, *tmp...)
			order(root.Left, &newTmp, res)
		}

		if root.Right != nil {
			newTmp := make([]string, 0, len(*tmp))
			newTmp = append(newTmp, *tmp...)
			order(root.Right, &newTmp, res)
		}
	}
	order(root, &[]string{}, &res)
	data := make([]string, len(res))
	for index, value := range res {
		data[index] = list2string(value)
	}
	return data
}

func list2string(data []string) string {
	return strings.Join(data, "->")
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	paths := binaryTreePaths(root)
	fmt.Println(paths)
}
