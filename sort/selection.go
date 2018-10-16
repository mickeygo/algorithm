package sort

// Selection 选择排序
type Selection struct {
	Elements []int
}

// Sort 选择排序
// 规则: 找到数组中最小的元素，与数组第一个值交换；再找剩余中最小的元素，与数组第二个值交换，以此类推
// 时间复杂度：(N-1)+(N-2)+(N-3)+...+2+1 = O(N2)
// 优点：简单易理解，且不占用额外的内存空间
// 缺点：不管是什么数组，都会重新排序一次。已排序的或是元素全相等的所用的排序时间一样长
func (s *Selection) Sort() []int {
	for i := 0; i < len(s.Elements)-1; i++ {
		index := 0           // 索引值，用于标识数组中最小元素的位置
		min := s.Elements[i] // 假设当前的值是最小的
		for j := i + 1; j < len(s.Elements); j++ {
			// 若内循环中有值小于当前的值，将该值设为最小值
			// 依次类推，直到内循环结束
			if min > s.Elements[j] {
				min = s.Elements[j]
				index = j
			}
		}

		// 若剩余数组中存在最小元素，将最小元素与当前值替换
		if index != 0 {
			s.Elements[index] = s.Elements[i]
			s.Elements[i] = min
		}
	}

	return s.Elements
}

func (s *Selection) Sort2() []int {

	return nil
}

func less(el []int, a int, b int) bool {
	return el[a] < el[b]
}

func exch(el []int, a int, b int) {
	temp := el[a]
	el[a] = el[b]
	el[b] = temp
}
