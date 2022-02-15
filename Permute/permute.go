package main

import "fmt"

//https://leetcode-cn.com/problems/permutations/
func permute(nums []int) [][]int {
	var result [][]int
	var backtracking func(i int, nums []int, track []int, used map[int]bool, result [][]int)
	backtracking = func(start int, nums []int, track []int, used map[int]bool, result [][]int)  {
		fmt.Println(track)
		if  len(track) == 3 {
			fmt.Println(track,444)
			p:=make([]int,len(track))
			copy(p,track)
			result = append(result, p)
			return
		}
		for i := start; i < len(nums); i++ {
			if used[i] == true {
				continue
			}
			track = append(track, nums[i])
			used[i] = true
			backtracking(i, nums, track, used, result)
			track = track[:len(track)-1]
			fmt.Println(track, 555)
			used[i] = false

		}

		//for _, num := range nums {
		//	if used[num] == true {
		//		continue
		//	}
		//	used[num] = true
		//	backtracking()
		//}

	}
	used := make(map[int]bool)
	fmt.Println(used)
	backtracking(0,nums, []int{}, used, result)
	return result
}
// 正确写法, 不能通过下标，来获取值，因为下边不能进行跳跃
func permuteV2(nums []int) [][]int {
	result := [][]int{}
	var backtracking func(nums []int, used map[int]bool, track []int, result *[][]int)

	backtracking = func(nums []int, used map[int]bool, track []int, result *[][]int) {
		if len(track) == 3 {
			tp := make([]int, len(track))
			copy(tp, track)
			// 注意用指针的方式，函数传值是复制传值
			*result = append(*result, tp)
			return
		}
		for _, num := range nums {
			if used[num] {
				continue
			}
			used[num] = true
			track = append(track, num)
			backtracking(nums, used, track, result)
			used[num] = false
			track = track[:len(track)-1]
		}
	}
	used := make(map[int]bool)
	backtracking(nums, used, []int{}, &result)
	return result
}
func main()  {
	//fmt.Println(permute([]int{1,2,3}))
	fmt.Println(permuteV2([]int{1,2,3}))
}


