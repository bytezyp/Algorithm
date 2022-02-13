package main

import "fmt"

var result [][]int
func combine(n, k int) [][]int {
	// 回溯法，不剪枝 or 剪枝
	if n <= 0 || k <= 0 || k > n {
		return nil
	}
	res := [][]int{}
	// track代表路径
	var backtracking func(n , k, startIndex int, track []int)
	backtracking = func(n, k, startIndex int, track []int) {
		// 回溯or递归终止条件
		// 遇到叶子结点或者本题中，len==k，认为可以终止
		// 为防止影响track本身，采用tmp临时变量操作加入路径
		if len(track) == k {
			tmp := make([]int, k)
			copy(tmp, track)
			res = append(res, tmp)
			//fmt.Println(track)
			//return
		}

		// 剪枝
		// for循环选择的起始位置之后的元素个数，已经不足我们需要的元素个数了，那么就没有必要搜索
		// n-startIndex+1 < k - len(track), 表示遍历的长度已经不足需要个数
		if len(track)+n-startIndex+1 < k {
			return
		}

		// startIndex即为每次添加的第一个数
		for i:=startIndex;i<=n;i++ {
			track = append(track, i) // 处理的节点
			backtracking(n,k,i+1,track) // 递归更深的
			track = track[:len(track)-1] // 撤销处理的节点
		}
	}
	backtracking(n,k,1, []int{})
	return res
}
func main()  {
	n, k := 4,2
	fmt.Println(combine(n ,k))
	//arr := []int{1, 2, 3}
	//temp := []int{}
	//copy(temp, arr)
	//fmt.Println(arr, temp)
}

