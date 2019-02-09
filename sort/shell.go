package sort

import (
	"fmt"
)

// Shell 希尔排序（分组插入排序）
type Shell struct {
}

// Sort 希尔排序
//
// 规则：将元素分组，每组内进行排序
// 时间复杂度：平均为 O(NlogN); 最差条件下为 O(N^2)
// 优点：
// 缺点：
func (s *Shell) Sort(arr []int) []int {
	if arr == nil || len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	N := len(arr)
	// 设置步长（增量）为 N/2, 后续以 1/2 速率递减
	for gap := N / 2; gap >= 1; gap = gap / 2 {
		// 插入排序
		// i 脚标，每组比较元素的最后的一个脚标
		// j 与 i 同一组元素的脚标
		for i := gap; i < N; i++ {
			// 对每个分组内的元素进行排序
			// 比较的阶长为 j-gap
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
		// 将每组的数据内部进行排序
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

// Sort3 排序
func (s *Shell) Sort3(arr []int) []int {
	N := len(arr)
	if N <= 1 {
		return arr
	}

	// 按初始值为 N/2 的间距进行分组，直至间距为 1 为止
	for gap := N / 2; gap >= 1; gap = gap / 2 {
		// 从第 gap 个元素开始，逐个对其组内的数据进行插入排序
		for i := gap; i < N; i++ {
			// 插入排序
			// 按增量距离 gap 进行插入排序
			for j := i - gap; j >= 0 && less(arr, j+gap, j); j -= gap {
				exch(arr, j, j+gap)
			}
		}
	}

	return arr
}
