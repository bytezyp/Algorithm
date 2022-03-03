package main

import "fmt"

// https://leetcode-cn.com/problems/valid-parentheses/
// 有效括号问题
func isValid(s string) bool {
	if len(s) <= 0 {
		return false
	}
	if len(s)%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 {
			// 说明字符串，还有值，但是栈没有值了；或者判断成对值不相等
			// 防止长度不相等的情况
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {

				return false
			}
			stack = stack[:len(stack)-1]
		}else{
			stack = append(stack, s[i])
		}
	}
	// 判断队列是否全部匹配删除
	if len(stack) > 0 {
		return false
	}
	return true
}

func main()  {
	s := "(]"
	fmt.Println(isValid(s))
}
