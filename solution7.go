package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head
	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			node.Next = l1
			node = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			node = l2
			l2 = l2.Next
		}
	}
	return head.Next
}
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	newList := head

	for l1 != nil || l2 != nil {
		if l1 == nil || (l2 != nil && l1.Val > l2.Val) {
			newList.Next = l2
			newList = l2
			l2 = l2.Next
		} else {
			newList.Next = l1
			newList = l1
			l1 = l1.Next
		}
	}
	return head.Next
}
