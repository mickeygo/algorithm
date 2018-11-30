package sort

// MergeSort 归并排序
//
// 思想：归并策略，分治策略
type MergeSort struct {
}

// SortUpDown 自顶向下的归并排序
func (s *MergeSort) SortUpDown(arr []int, lo, hi int) {

}

// SortDownUp 自底向上的归并排序
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
	if i < mid {
		for i < mid {
			temp[k] = arr[i]
			k++
			i++
		}
	}
	if j < hi {
		for j < hi {
			temp[k] = arr[j]
			k++
			j++
		}
	}
}
