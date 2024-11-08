package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListNode(node *ListNode) *ListNode {
	if node.Next == nil {
		if node != nil {
			fmt.Println("return head", node.Val)
		}
		return node
	}
	fmt.Println("visis next node", node.Next.Val)
	newhead := reverseList(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newhead
}

func reverseList(head *ListNode) *ListNode {
	return reverseListNode(head)
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	reverseList(head)
}
