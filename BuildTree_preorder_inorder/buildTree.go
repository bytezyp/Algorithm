package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 105 根据前序遍历和中序遍历，构成树形
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	left := findRootIndex(preorder[0], inorder)
	root := &TreeNode{
		Val: preorder[0],
		Left: buildTree(preorder[1:left+1], inorder[:left]),
		Right: buildTree(preorder[left+1:], inorder[left+1:]),
	}
	return root
}

func findRootIndex(target int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			return i
		}
	}
	return -1
}

func main()  {

}

