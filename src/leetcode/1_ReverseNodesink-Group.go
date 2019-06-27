package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
 }

// 一段一段反转，然后连接
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {// 判断输入参数
		return nil
	}
	if k == 1 {
		return head
	}
	nodes := make([]ListNode, k)
	headCopy := head
	for i := 0; i < k; i++ {
		if head != nil {
			nodes[i] = *head
			head = head.Next
		} else {
			return headCopy
		}
	}
	headNew := reverseListNode(nodes, k) //反转一段
	tailHeadNew := tailOfList(headNew)
	tailHeadNew.Next = reverseKGroup(head, k)// 连接后一段
	return headNew
}

func reverseListNode(nodes []ListNode, len int) *ListNode {
	if len == 1 {
		return &nodes[0]
	}
	for i := len - 2; i >= -1; i-- {
		if i >= 0 {
			nodes[i+1].Next = &nodes[i]
		} else {
			nodes[i+1].Next = nil
		}
	}
	return &nodes[len - 1]
}

func tailOfList(head *ListNode) *ListNode {
	for {
		if head.Next == nil {
			break;
		}
		head = head.Next
	}
	return head
}

func main() {
	node1 := ListNode{1,nil}
	node2 := ListNode{2,nil}
	node3 := ListNode{3,nil}
	node4 := ListNode{4,nil}
	node5 := ListNode{5,nil}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	nodeNew := reverseKGroup(&node1, 2)
	fmt.Println(*nodeNew)
}
