package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
// 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dump := &ListNode{0,nil}
	cur := dump
	for list1 != nil && list2 != nil {
		//fmt.Println(list1.Val, list2.Val)
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		}else{
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 == nil {
		cur.Next = list2
	}
	if list2 == nil {
		cur.Next = list1
	}
	return dump.Next

}

func main()  {
	list1 := &ListNode{1,&ListNode{3,&ListNode{5, nil}}}

	list2 := &ListNode{1,&ListNode{2,&ListNode{8,nil}}}

	list := mergeTwoLists(list1,list2)

	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}



