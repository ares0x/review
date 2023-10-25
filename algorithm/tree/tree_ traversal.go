package algorithm

import "fmt"

type TreeNode struct{
	Val int
	Left, Right *TreeNode
}

// preOrderTraversal 前序 根 - 左 - 右
func preOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrderTraversal(root.Left)
	preOrderTraversal(root.Right)
}

// inOrderTraversal 中序 左 - 根 -  右
func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left)
	fmt.Println(root.Val)
	inOrderTraversal(root.Right)
}

// postOrderTraversal 后序遍历 左 - 右 - 根
func postOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postOrderTraversal(root.Left)
	postOrderTraversal(root.Right)
	fmt.Println(root.Val)
}

// levelOrderTraversal 层序遍历
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    result := [][]int{}
    nodeArray := []*TreeNode{root}
    for len(nodeArray) != 0 {
        res := []int{}
        size := len(nodeArray)
        for i := 0; i < size; i++ {
            node := nodeArray[0]
            res = append(res, node.Val)
            nodeArray = nodeArray[1:]

			if node.Left != nil {
				nodeArray = append(nodeArray, node.Left)
			}
			if node.Right != nil {
				nodeArray = append(nodeArray, node.Right)
			}
        } 
        result = append(result,res)
    }
    return result
}
