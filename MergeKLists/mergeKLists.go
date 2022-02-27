package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists)
}

func merge(lists []*ListNode) *ListNode {
	if len(lists) < 2 {
		return lists[0]
	}
	mid := len(lists)/2
	left := merge(lists[:mid])
	right := merge(lists[mid:])
	return mergeTwoList(left, right)
}
// 合并两个有序链表
func mergeTwoList(list1, list2 *ListNode) *ListNode {
	dump := &ListNode{}
	cur := dump
	for list1 != nil && list2 != nil {
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
	arr := []*ListNode{
		&ListNode{1,&ListNode{2,&ListNode{4,nil}}},
		&ListNode{4,&ListNode{6,&ListNode{7,nil}}},
		&ListNode{3,&ListNode{8,&ListNode{9,&ListNode{11,nil}}}},
		&ListNode{2,&ListNode{8,&ListNode{9,&ListNode{11,nil}}}},
	}
	list := mergeKLists(arr)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
