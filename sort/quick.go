package sort

// QuickSort 快速排序
type QuickSort struct{}

// Sort 快速排序
//
// 思想：
// 时间复杂度：
//	最优时间:O(nlogn)；最坏情况下：O(n^2)
// 空间复杂度
// 稳定性：
func (s *QuickSort) Sort(arr []int, lo, hi int) {
	pivot := 0
	if lo < hi {
		pivot = s.partition(arr, lo, hi)
		s.Sort(arr, lo, pivot-1)
		s.Sort(arr, pivot+1, hi)
	}
}

func (s *QuickSort) partition(arr []int, lo, hi int) int {
	pivot := arr[lo]

	return pivot
}
