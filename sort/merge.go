package sort

// MergeSort 归并排序
//
// 思想：归并策略，分治策略
// 时间复杂度: O(nlogn)
// 空间复杂度: O(n)
type MergeSort struct {
}

// SortUpDown 自顶向下的归并排序(基于递归的归并排序)
func (s *MergeSort) SortUpDown(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := s.SortUpDown(arr[:mid])
	right := s.SortUpDown(arr[mid:])
	return s.merge3(left, right)
}

// SortDownUp 自底向上的归并排序(基于迭代的归并排序)
// 两路归并思想：
//	将 N 数组看成是 N 个长度为 1 的有序序列
//	将这个 N 个有序序列两两归并，得到 N/2 个长度为 2 的有序序列
//	重复上面一步，直到得到 1 个长度为 N 的有序序列为止
func (s *MergeSort) SortDownUp(arr []int, lo, hi int) {
}

// 原地归并的抽象方法
func (s *MergeSort) merge1(arr []int, lo int, mid int, hi int) {
	aux := make([]int, len(arr))
	for i, item := range arr {
		aux[i] = item
	}

	// lo 低位；hi 高位; mid = (lo+hi)/2
	i, j := lo, mid+1 // i,j 分别为指向以 mid 分割的两数组低
	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > hi {
			arr[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			arr[k] = aux[j]
			j++
		} else {
			arr[k] = aux[i]
			i++
		}
	}
}

// 原地归并的抽象方法 2
func (s *MergeSort) merge2(arr []int, lo int, mid int, hi int) {
	temp := make([]int, len(arr))
	k, i, j := 0, lo, mid+1 // i,j 分别为指向以 mid 分割的两数组低位
	for i <= mid && j <= hi {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}

		k++
	}

	// 将后续没循环完的数组值取出赋值
	for i < mid {
		temp[k] = arr[i]
		k++
		i++
	}

	for j < hi {
		temp[k] = arr[j]
		k++
		j++
	}
}

// merge3 巧妙利用 go 语言特性, 效果同 merge2
func (s *MergeSort) merge3(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return
}
