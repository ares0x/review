package algorithm

type TreeNode struct{
	Val int
	Left, Right *TreeNode
}

func inOrder(root *TreeNode) []int{
	res := []int{}
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode){
		if node == nil {
			return
		}
		inOrder(node.Left)
		res = append(res, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	return res
}