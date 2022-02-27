package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 前序遍历 根左右
func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil  {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return result
}

func main()  {
	
}


