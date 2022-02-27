package main

import (
	"fmt"
	"math"
)

// 给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
func divide(a, b int) int {
	// 无线小的数 直接返回
	if a == math.MinInt32 && b == -1 {
		return math.MinInt32
	}
	res := 0
	// 除数绝对值最大，结果必然是0或者1
	if b == math.MinInt32 {
		if a == b {
			return 1
		}else{
			return 0
		}
	}
	// 被除数先减去一个除数
	if a == math.MinInt32 {
		a -= -abs(b)
		res += 1
	}
	sign := 1
	if ( a > 0 && b < 0 ) || (a < 0 && b > 0){
		sign = -1
	}
	a = abs(a)
	b = abs(b)
	for i := 31; i >= 0; i-- {
		if (a >> i) - b >= 0 { // a >= (b << i)
			a = a - (b << i)
			if res > math.MaxInt32 - (1 << i) {
				return math.MinInt32
			}
			res += 1<< i
		}
	}
	return sign * res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main()  {
	 //fmt.Println(3*1 <<3)
	 fmt.Println(divide(5,2))

	 s1 := make([]int, 0, 10)
	 var appenFunc = func(s []int) {
		 s = append(s ,10,20,30)
		 fmt.Printf("%p %v   %p\n", &s, s, &s[0])
	 }
	fmt.Printf("%p %v   %p\n", &s1, s1, &s1)
	fmt.Println(s1)
	appenFunc(s1)
	 fmt.Println(s1)

	 fmt.Println(s1[:10])
	 fmt.Println(s1[:])
}


