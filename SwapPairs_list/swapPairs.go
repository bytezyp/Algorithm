package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
// 链表两两交换
func swapPairs(head *ListNode) *ListNode {
	dump := &ListNode{0 ,head}
	cur := dump
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next
		node1.Next = node2.Next
		cur.Next = node2
		node2.Next = node1
		cur = node1
	}
	return dump.Next
}

func main()  {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4,&ListNode{5,nil}}}}}
	//for l != nil {
	//	fmt.Println(l.Val)
	//	l = l.Next
	//}
	pre := swapPairs(l)
	for pre != nil {
		fmt.Println(pre.Val)
		pre = pre.Next
	}
}
