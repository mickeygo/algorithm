package queue

import (
	"errors"
	"sync"
)

// CilcleQueue 基于数组的环形队列
//
// 使用环形队列, 可减少数据在数组中的移动
type CilcleQueue struct {
	items []interface{}
	front uint // 队头指针
	rear  uint // 队尾指针，注意：队尾指针指向的数组中尾部元素的后一位(初始化时都指向同一位置)
	// 队列容量
	Capacity uint

	mutex sync.Mutex
}

// NewCilcleQueue 新建一个基于数组的环形队列
//
// N 队列的容量
func NewCilcleQueue(N uint) *CilcleQueue {
	return &CilcleQueue{
		items:    make([]interface{}, 0, N),
		front:    0,
		rear:     0, // 初始化时队尾指针指向 0 位置
		Capacity: N,
	}
}

// Push 向队列加入数据
func (queue *CilcleQueue) Push(item interface{}) (bool, error) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	// 检查队列是否已满
	// 当尾部差一步就追上头部时，表示队列已满
	if (queue.rear+1)%queue.Capacity == queue.front {
		return false, errors.New("Queue is fully")
	}

	// 在尾部加入元素
	queue.items[queue.rear] = item
	// 尾部指针向后移动
	queue.rear = (queue.rear + 1) % queue.Capacity

	return true, nil
}

// Pop 出列
//
// 当队列中没有数据时，返回 nil
func (queue *CilcleQueue) Pop() (bool, interface{}) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	// 检查队列中是否有数据
	if queue.front == queue.rear {
		return false, errors.New("Queue is empty")
	}

	// 获取 head 元素
	// 注：这里并没有将 head 出元素设置为 nil
	item := queue.items[queue.front]

	// 移动头指针
	queue.front = (queue.front + 1) % queue.Capacity

	return true, item
}

// Count 队列中元素的数量
func (queue *CilcleQueue) Count() uint {
	return (queue.rear - queue.front + queue.Capacity) % queue.Capacity
}
