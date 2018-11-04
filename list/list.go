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
func CheckCycle(head *Node) (bool, *Node) {
	p1 := head
	p2 := head

	// 检查快指针，直到指向链表结尾，或与慢指针重叠
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		// 若两节点碰撞, 退出循环
		if p1 == p2 {
			break
		}
	}

	// 此时链表中没有环
	if p2 == nil || p2.Next == nil {
		return false, nil
	}

	/*将slow指向首部,fast指向碰撞处,两者距离环路
	 *起始处k步,若两者以相同速度移动,则必定会在环
	 *路起始处碰在一起*/
	p1 = head
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return true, p2
}

type NodeRune struct {
	Data rune
	Next *NodeRune
}

// UnionSorted 合并两有序的链表
//
// 返回一个有序的链表
// 思想：1. 先找到两者中较小的头
//		2. 定义两哨兵，依次比较，直至到达一方的尾部
//		3. 将后续的元素移至链表尾部
// 注意点：此时间复杂度为 O(N), 关键点是依次
func UnionSorted(head1 *NodeRune, head2 *NodeRune) *NodeRune {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var head, p *NodeRune
	p1, p2 := head1, head2

	// 找到起点小的节点
	if head1.Data < head2.Data {
		p = head1
		p1 = p1.Next
	} else {
		p = head2
		p2 = p2.Next
	}

	head = p

	// 两链表元素比较大小，直到有一方为 nil
	// 巧妙之处，而非使用内嵌循环
	for p1 != nil && p2 != nil {
		if p1.Data <= p2.Data {
			p.Next = p1
			p = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p = p2
			p2 = p2.Next
		}
	}

	// 将 p1 后续元素直接移到链表中
	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	return head
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
