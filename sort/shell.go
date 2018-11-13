package sort

import (
	"fmt"
)

// Shell 希尔排序（分组插入排序）
type Shell struct {
}

// Sort 希尔排序
//
// 规则：
// 时间复杂度：
// 优点：
// 缺点
func (s *Shell) Sort(arr []int) []int {
	if arr == nil || len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	N := len(arr)
	// 设置步长为 N/2, 后续以 1/2 速率递减
	for gap := N / 2; gap >= 1; gap = gap / 2 {
		// 插入排序
		// i 脚标，每组比较元素的最后的一个脚标
		// j 与 i 同一组元素的脚标
		for i := gap; i < N; i++ {
			// 对每个分组内的元素进行排序
			for j := i; j >= gap && less(arr, j, j-gap); j -= gap {
				exch(arr, j, j-gap)
			}
		}
	}

	return arr
}

// Sort2 希尔排序
func (s *Shell) Sort2(arr []int) []int {
	if arr == nil || len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	h, N := 1, len(arr)
	// 设置最初的最大步长
	for h < N/3 {
		h = h*3 + 1
	}
	// 按步长分组进行比较，直至步长为 1 为止
	for h >= 1 {
		// 将数组变为 h 有序
		for i := h; i < N; i++ {
			// 将 arr[i] 插入到 arr[i-h], arr[i-2*h], arr[i-3*h]... 之中
			// 从第i个元素开始，依次次和前面已经排好序的i-h个元素比较，如果小于，则交换
			for j := i; j >= h && less(arr, j, j-h); j -= h {
				exch(arr, j, j-h)
			}
		}

		fmt.Println(arr)
		// 步长按三分之一递减
		h = h / 3
	}

	return arr
}
