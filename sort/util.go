package sort

// less 比较数组中索引 a、b 的值的大小
// 当 [a] < [b], 返回 1
// 当 [a] > [b], 返回 -1
// 当 [a] = [b], 返回 0
func compare(el []int, a int, b int) int {
	if el[a] < el[b] {
		return 1
	} else if el[a] > el[b] {
		return -1
	} else {
		return 0
	}
}

// less 比较数组中索引 a、b 的值的大小, 当 [a] < [b] 时返回 true
func less(el []int, a int, b int) bool {
	return el[a] < el[b]
}

// exch 交换 a、b 在数组 el 中的值
func exch(el []int, a int, b int) {
	el[a], el[b] = el[b], el[a]
}
