package algorithm

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//  BFS
//  队列
func levelOrder(root *TreeNode)[][]int  {
	if root == nil {
		return nil
	}
	result := [][]int{}
	tmp := []*TreeNode{root}
	for i := 0; i< len(tmp); i++ {
		result = append(result, )
	}
}