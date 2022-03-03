package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 后续遍历 左右中
func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			fmt.Println(root.Val)
			result = append(result, root.Val)
			prev = root
			root = nil
		}else{
			stack = append(stack, root)
			root = root.Right
		}
	}
	return result
}

func main()  {
	root := &TreeNode{3, &TreeNode{9, nil, nil},&TreeNode{4,&TreeNode{5,nil,nil}, &TreeNode{7,nil,nil}}}
	fmt.Println(postorderTraversal(root))
}
