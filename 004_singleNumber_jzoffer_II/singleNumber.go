package main

import "fmt"

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums{
			total += int32(num) >> i &1
			//fmt.Println(total)
		}
		//fmt.Println(total)
		// | 二进制 求和
		if  total % 3 > 0 {
			ans |= (1<<i)
			fmt.Printf("%08b\n", ans)
		}
	}
	//fmt.Println(ans)
	return int(ans)
}

func main()  {
	fmt.Println(singleNumber([]int{2,2,2,9}))
	//a := 1<<3
	//b := 3|a
	//fmt.Println(b, a)
	//fmt.Println(1<<3|3)
	//fmt.Println(5|1<<3)
	//fmt.Println(1&1)
	//fmt.Println(1&1)
	//fmt.Println(0|1)
	//fmt.Println(6|1)
}




