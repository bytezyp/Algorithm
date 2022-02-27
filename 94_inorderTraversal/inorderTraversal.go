package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 中序遍历，左中右
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := []int{}
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		result = append(result, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right //  判断最后一个节点，是否有右节点
	}
	return result
}

func main()  {
	
}


