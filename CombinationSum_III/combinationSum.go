package main

import "fmt"

// 组合总和 III
func combinationSum(k, n int) [][]int {
	result := [][]int{}
	//max := 9
	var backtracking func(k int, n int, start int, sum int, track []int)
	backtracking = func(k int, n int, start int, sum int, track []int) {
		if sum > n || len(track) > k{
			return
		}
		if  sum == n && len(track) == k {
			tmp := make([]int, k)
			copy(tmp, track)
			result = append(result, tmp)
			return
		}
		for i := start; i <= 9-(k-len(track))+1; i++ {
			sum = sum + i
			track = append(track, i)
			backtracking(k, n, i + 1, sum, track)
			sum = sum - i
			track = track[:len(track)-1]
		}
	}
	backtracking(k, n,1, 0, []int{})
	return result
}

func main()  {
	fmt.Println(combinationSum(3,9))
}
