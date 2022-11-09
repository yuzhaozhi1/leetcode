package main

type Node struct {
	Val      int
	Children []*Node
}

func max1(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// maxDepthN
//  @Description: 559. N 叉树的最大深度
//  @param root:
//  @return int:
func maxDepthN(root *Node) int {
	depth := 0
	var order func(root *Node, level int)

	order = func(root *Node, level int) {
		if root == nil {
			return
		}
		depth = max1(depth, level)
		for _, child := range root.Children {
			order(child, level+1)
		}
	}
	order(root, 1)
	return depth
}
