package main

func reverseList(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := reverseList(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

// k个长度一组逆转链表 当长度不足k时也会逆转
func reverseByKLength(node *ListNode, k int) *ListNode {
	// 当前逆转区间的头和尾
	start, end := node, node
	for i := 0; i < k-1; i++ {
		if end == nil {
			continue
		}
		end = end.Next
	}
	// 说明长度不足k或者已经到最后一组 直接将逆转的头结点返回
	if end == nil || end.Next == nil {
		return reverseList(start)
	}
	// 递归调用逆转后面所有并返回头节点，连接到当前尾结点之后
	newTail := reverseList(end.Next)
	// 断开链表
	end.Next = nil
	// 当前区间逆转后的链表头
	curHead := reverseList(start)
	// start在逆转后变成尾结点连接两段
	start.Next = newTail
	// 返回从start开始的部分递归的结果
	return curHead
}

func swapPairs(head *ListNode) *ListNode {
	return reverseByKLength(head, 4)
}
