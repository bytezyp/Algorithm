package main

import "fmt"

// 有些问题
func combinationSum(k, n int) [][]int {
	result := [][]int{}
	//max := 9
	var backtracking func(k int, n int, start int, sum int, track []int)
	backtracking = func(k int, n int, start int, sum int, track []int) {
		if sum > n || len(track) > k{
			return
		}
		if sum == n && len(track) == k {
			// 用于检验是否有重复的
			m := map[int]struct{}{}
			for _, v := range track {
				m[v] = struct{}{}
			}
			if len(m) != k {
				return
			}
			tmp := make([]int, k)
			copy(tmp, track)
			result = append(result, tmp)
		}
		for i := start; i <= 9-(k-len(track))+1; i++ {
			sum = sum + i
			track = append(track, i)
			backtracking(k, n, start + 1, sum, track)
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
