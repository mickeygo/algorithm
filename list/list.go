package list

import (
	"fmt"
)

// Node 单链表结构
type Node struct {
	Data interface{}
	Next *Node
}

// List 链表结构
type List struct {
	size uint64
	head *Node // 头
	tail *Node // 尾
}

// Print 打印节点元素
//
// head 指向链表头节点的元素
func Print(head *Node) {
	p := head
	if p == nil {
		fmt.Printf("node is empty ... \n")
		return
	}

	for p != nil {
		fmt.Printf("%v ... \n", p.Data)

		p = p.Next
	}
}

// Reverse 链表反转
//
// head 链表的起始节点
func Reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	p := head // p 用于记录当前节点
	next := head.Next

	p.Next = nil // 将 head 节点的 next 指向 nil

	// 从第二个节点开始
	for next != nil {
		pNext := next.Next // 临时获取下一个节点

		next.Next = p
		p = next

		next = pNext
	}

	return next
}
