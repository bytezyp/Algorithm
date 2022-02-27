package main

import "fmt"

type uint64Slice []uint64

func quickSort(nums uint64Slice, start, end int)  {
	var mid int
	tmpStart := start
	tmpEnd := end
	//递归结束条件
	if start >= tmpEnd {
		return
	}
	pivot := nums[start]
	for start < end {
		for start < end && nums[end] > pivot{
			end--
		}
		if start < end {
			nums.swap(start, end)
			start++
		}

		for start < end && nums[start] < pivot {
			start++
		}
		if start < end {
			nums.swap(start,end)
			end--
		}
	}
	nums[start] = pivot
	mid = start
	quickSort(nums, tmpStart, mid-1)
	quickSort(nums, mid+1, tmpEnd)
}

func (num uint64Slice) swap(i,j int)  {
	num[i], num[j] = num[j], num[i]
}
func main()  {
	numbers := []uint64{5,4,20,3,8,2,8}
	quickSort(numbers,0,len(numbers)-1)
	nn := 5/2
	fmt.Println(nn)
	fmt.Println(numbers)
}



