package main

import "fmt"

//leetcode 组合求和 39 题
// https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int  {
	var result [][]int
	var backtracking func(start int, candidates []int, sum int, target int, trace []int)
	backtracking = func(start int, candidates []int, sum int, target int, trace []int) {
		//fmt.Println(trace, sum, target)
		if sum > target{
			return
		}
		if sum == target {
			tmp := make([]int, len(trace), len(trace))
			copy(tmp, trace)
			result = append(result, tmp)
			return
		}
		for i := start; i <= len(candidates)-1; i++ {
			sum += candidates[i]
			trace = append(trace, candidates[i])
			backtracking(i, candidates, sum, target, trace)
			sum -= candidates[i]
			trace = trace[:len(trace)-1]
		}
	}
	backtracking(0,candidates, 0, target,[]int{})
	return result
}

func main()  {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}
