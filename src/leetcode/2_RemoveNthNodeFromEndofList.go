package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	headCopy := head
	len := lenOfList(head)
	if (len == 1 && n == 1) || (n > len) {
		return nil
	}
	if len == n {
		head = head.Next
		return head
	}
	for i := 0; i < len - n - 1; i++ {//需要保留的前一段最后一个节点
		if headCopy != nil {
			headCopy = headCopy.Next
		} else {
			return head
		}
	}
	// 需要删除的节点
	if headCopy != nil {
		frontEndNode := headCopy
		deleteNode := headCopy.Next
		if deleteNode == nil || deleteNode.Next == nil {
			frontEndNode.Next = nil
		} else {
			frontEndNode.Next = deleteNode.Next
		}
	}
	return head
}

func lenOfList(head *ListNode) int {
	numNode := 0
	for {
		if head == nil {
			break
		} else {
			numNode++
			head = head.Next
		}
	}
	return numNode
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
	nodeNew := removeNthFromEnd(&node1, 2)
	fmt.Println(*nodeNew)
}