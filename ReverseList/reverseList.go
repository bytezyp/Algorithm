package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func main()  {
	list := &ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4,nil}}}}
	//for list != nil {
	//	fmt.Println(list.Val)
	//	list = list.Next
	//}
	ll := reverseList(list)
	for ll != nil {
		fmt.Println(ll.Val)
		ll = ll.Next
	}
}

