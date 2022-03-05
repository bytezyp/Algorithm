package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dump := &ListNode{0, head}
	cur := dump
	for cur.Next != nil {
		left := cur
		for i := 0; i < k; i++ {
			cur = cur.Next
			if  cur == nil {
				return dump.Next
			}
		}
		//fmt.Println(left.Next.Val, cur.Val, 22)
		newHead, newTail := reverse(left.Next, cur)
		fmt.Println(newHead.Val, newTail.Val, 33)
		left.Next = newHead
		cur = newTail
	}
	return dump.Next
}

func reverse(left, right *ListNode) (top,tail *ListNode) {
	rNext := right.Next
	//dump := &ListNode{0,left}
	//
	//cur := dump
	//for cur.Next != rNext {
	//	node1 := cur.Next
	//	node2 := cur.Next.Next
	//	node1.Next = node2.Next
	//	node2.Next = node1
	//	cur.Next = node2
	//	cur = node1
	//}
	cur := left
	for cur != right {
		nex := cur.Next
		cur.Next = rNext
		rNext = cur
		cur = nex
	}
	//fmt.Println(dump.Next.Val, cur.Val,dump.Next.Next.Val)
	//cur.Next = rNext
	//fmt.Println(cur.Next.Val, cur.Next.Next.Val)
	//panic(11)
	return right, left
}

func main()  {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4,&ListNode{5,&ListNode{6,nil}}}}}}
	pre := reverseKGroup(l, 3)
	fmt.Println(pre.Val, pre.Next.Val, 66)
	for pre != nil {
		fmt.Println(pre.Val)
		pre = pre.Next
	}
}


