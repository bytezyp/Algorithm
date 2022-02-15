package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var result [][]int
	var backtracking func(nums []int, start int, used []bool, track []int, result *[][]int)
	// 排序是在一样的数字在一起，否则之前的动态，不能合并。
	sort.Ints(nums)
	size := len(nums)
	backtracking = func(nums []int, start int, used []bool, track []int, result *[][]int) {
		fmt.Println(track)
		if len(track) == size {
			tp := make([]int, size)
			copy(tp, track)
			*result = append(*result, track)
			return
		}

		for i := start; i < len(nums); i++ {
			if i - 1 >= 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			// 用于再次过滤的时候，因为每次都是从头开始
			if used[i] {
				continue
			}
			used[i] = true
			track = append(track, nums[i])
			backtracking(nums, start, used, track, result)
			used[i] = false
			track = track[0:len(track)-1]
		}
	}
	used := make([]bool, 3)
	backtracking(nums, 0, used, []int{}, &result)
	return result
}

func main()  {
	fmt.Println(permuteUnique([]int{1,1,2}))
}



