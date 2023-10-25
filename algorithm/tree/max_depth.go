package algorithm

type TreeNode struct{
	Val int
	Left,Right *TreeNode
}

// DFS 
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
        return left + 1
    }
    return right + 1
}

func maxDepthBack(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	level := 0
	for len(stack) != 0 {
	  size := len(stack)
	  for i := 0; i < size; i++ {
		  node := stack[0]
		  stack = stack[1:]
		  if node.Left != nil {
			 stack = append(stack, node.Left)
		  }
		  if node.Right != nil {
			 stack = append(stack, node.Right)
		  }
	  }
	  level++
	}
	return level
 }