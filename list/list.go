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

// Delete 删除链表的第 N 个元素
//
// 返回重新指向的链表头地址
func Delete(head *Node, N uint) *Node {
	if head == nil {
		return nil
	}

	// 删除头元素
	if N == 0 {
		return head.Next
	}

	var count uint
	var prev *Node
	p := head
	for p != nil && count < N {
		prev = p
		count++
		p = p.Next
	}

	// 当 N 不超出链表个数据时，将前一元素的 next 指向后一元素
	if p != nil {
		prev.Next = p.Next
	}

	return head
}

// Search 查找链表中第 N 个元素(索引从 0 开始)的值
//
// 返回第 N 个元素指针
// 若链表数目小于 N 会返回 nil
func Search(head *Node, N uint) *Node {
	var count uint
	p := head
	for p != nil && count < N {
		count++
		p = p.Next
	}

	return p
}

// Reverse 链表反转
//
// head 链表的起始节点
// 返回 node 的尾部
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

	return p
}

// CheckCycle 检查链表中是否有环路
//
// 返回是否有环路，以及进入环路的节点索引（从0开始）
// 现象：链表尾部指向了链表链路本身
// 方法：https://www.aliyun.com/jiaocheng/1444871.html
func CheckCycle(head *Node) (bool, uint) {
	p1 := head
	p2 := head

	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		// 若两节点碰撞
		if p1 == p2 {
			break
		}
	}

	return false, 0
}

// UnionSorted 合并两有序的链表
//
// 返回一个有序的链表
func UnionSorted(sortedhead1 *Node, sortedhead2 *Node) *Node {

	return nil
}

// IndexOfMiddle 检索链表的中间节点
func IndexOfMiddle(head *Node) *Node {
	p1 := head
	p2 := head

	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return p1
}
