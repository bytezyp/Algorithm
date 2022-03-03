package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var ret int
	var dfs func(root *TreeNode, k int)

	dfs = func(root *TreeNode, k int)  {
		if root == nil {
			return
		}
		dfs(root.Right, k)
		if k -= 1; k == 0 {
			ret = root.Val
			return
		}
		dfs(root.Left, k)
	}
	 dfs(root, k)
	return ret
}
func main()  {
	
}


