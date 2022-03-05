package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dump := &ListNode{-1, head}
	cur := dump
	var pre *ListNode
	for i := 0; i < left; i++ {
		pre = cur
		cur = cur.Next
	}
	curLeft := cur
	for i := 0; i < right - left; i++ {
		cur = cur.Next
	}
	curRight := cur
	newLeft, _ := reverse(curLeft, curRight)
	pre.Next = newLeft
	return dump.Next
}
func reverse(left, right *ListNode) (newRight *ListNode, newLeft *ListNode) {
	rNext := right.Next
	stop := rNext
	cur := left
	for cur != stop {
		next := cur.Next
		cur.Next = rNext
		rNext = cur
		cur = next
	}
	return right,left
}
func main()  {
	list := &ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4,&ListNode{5,nil}}}}}
	//for list != nil {
	//	fmt.Println(list.Val)
	//	list = list.Next
	//}
	ll := reverseBetween(list, 1,5)
	for ll != nil {
		fmt.Println(ll.Val)
		ll = ll.Next
	}
}
