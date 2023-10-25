package algorithm

// leetcode 226 翻转二叉树
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归
/*
	递归：
	1. 确定递归函数和返回值
	2. 确定终止条件
	3. 确定单层递归逻辑
**/
func invertTree(root *TreeNode) *TreeNode{
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}

// 迭代

