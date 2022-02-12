package main

type ListNode struct {
	Val int
	Next *ListNode
}
func detectCycle(head *ListNode) *ListNode {
	slow,fast := head,head
	for slow != nil {
		slow = slow.Next
		
	}
}
