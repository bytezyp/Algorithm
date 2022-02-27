package main

import "fmt"

type uint64Slice []uint64

// 堆排序
func sortHeap(nums uint64Slice)  {
	length := len(nums)
	buildMaxHeap(nums, length)
	for i := length-1; i > 0; i-- {
		nums.swap(0, i)
		length -= 1
		heapify(nums, 0, length)
	}
}
// 构建大顶堆
func buildMaxHeap(nums uint64Slice, length int)  {
	for i := length/2; i >= 0; i-- {
		fmt.Println(i, nums)
		heapify(nums, i, length)
		fmt.Println(nums)
	}
}
// 堆调整
func heapify(nums uint64Slice, i, length int)  {
	left := 2*i + 1 //  i = 2 left = 5 right = 5
	right := 2*i + 2
	largest := i
	if left < length && nums[left] > nums[largest]{
		largest = left
	}

	if right < length && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums.swap(i, largest)
		heapify(nums, largest, length)
	}
}

func (num uint64Slice) swap(i, j int)  {
	num[i], num[j] = num[j], num[i]
}
func main()  {
		nn := []uint64{5,2,7,3,6,1,4}
		sortHeap(nn)
		fmt.Println(nn)
}







