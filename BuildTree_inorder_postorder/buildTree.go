package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 106. 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	// 先找根节点
	nodeValue := postorder[len(postorder)-1]
	// 从中序遍历中找到一分为二的点，左边为左子树，右边为右子树
	left := findRootIndex(inorder, nodeValue)
	root := &TreeNode{
			Val: nodeValue,
			Left: buildTree(inorder[:left], postorder[:left]),
			Right: buildTree(inorder[left+1:],postorder[left:len(postorder)-1])}
	return root
}
func findRootIndex(inorder []int, target int) (index int)  {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			return i
		}
	}
	return -1
}
func main()  {
	inorder := []int{9,3,15,20,7}
	postorder := []int{9,15,7,20,3}
	buildTree(inorder, postorder)
}
