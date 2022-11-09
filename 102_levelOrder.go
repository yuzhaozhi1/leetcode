package main

import "container/list"

// 102. 二 叉树的层序遍历

// levelOrder1
//  @Description:递归解法
//  @param root:
//  @return [][]int:
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	var level int

	var order func(root *TreeNode, level int)
	order = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(res) == level {
			res = append(res, make([]int, 0, 10))
		}
		res[level] = append(res[level], root.Val)
		if root.Left != nil {
			order(root.Left, level+1)
		}
		if root.Right != nil {
			order(root.Right, level+1)
		}
	}
	order(root, level)
	return res
}

// levelOrder
//  @Description: 二叉树的层序遍历
//  @param root:
//  @return [][]int:
func levelOrder(root *TreeNode) [][]int {
	var arr [][]int
	if root == nil {
		return arr
	}
	// 构建一个链表
	li := list.New()
	li.PushBack(root) // 把root 节点添加到链表的结尾
	var tmpArr []int

	// 原理就是利用队列的先进先出的特点
	// 第二层就是两个数据入队遍历那两个数据
	// 第三层就是4个数据入队,然后遍历那4个
	// 直到最后一层后,队列中没有数据了, 此时最外层的li.Len() 的长度为0, 循环结束
	for li.Len() > 0 {
		// 保存当前层的长度
		length := li.Len()
		for i := 0; i < length; i++ {
			// 返回第一个元素
			node := li.Remove(li.Front()).(*TreeNode)
			tmpArr = append(tmpArr, node.Val)
			if node.Left != nil {
				li.PushBack(node.Left)
			}
			if node.Right != nil {
				li.PushBack(node.Right)
			}
		}
		arr = append(arr, tmpArr) // 将层数据添加到结果arr 中
		tmpArr = []int{}          // 将数据重置为空
	}
	return arr
}
