package algorithm

// BinarySearch 二分查找，算法时间度 O(logN)
// arrs 被检索的数组，注：必须为已排序的数组
// key 表示要查找的值
// return 索引位置，若没找到，返回 -1
func BinarySearch(arrs []int, key int) (int, count int) {
	lo := 0             // 低位
	ho := len(arrs) - 1 // 高位
	for lo <= ho {
		count++

		mid := lo + (ho-lo)/2
		if key < arrs[mid] {
			ho = mid - 1 // 将高位移至原 mid 的左边一位(因为 mid 不满足条件，再次查询时应排除)
		} else if key > arrs[mid] {
			lo = mid + 1 // 将低位移至原 mid 的右边一位
		} else {
			return mid, count
		}
	}

	return -1, count
}
