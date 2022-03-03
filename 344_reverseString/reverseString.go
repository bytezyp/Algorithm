package main

import "fmt"

func reverseString(s string)  {
	ss := []byte(s)
	for i := 0; i <= (len(ss)-1)/2; i++ {
		ss[i], ss[len(ss)-1-i] = ss[len(ss)-1-i],ss[i]
	}
	fmt.Println(string(ss))
	return
}

func reverseStringV1(s []byte)  {
	length := len(s)-1
	for left , right := 0, length; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
	fmt.Println(string(s))
}

func main()  {
	//reverseString("12345")
	reverseStringV1([]byte("56789"))
}



