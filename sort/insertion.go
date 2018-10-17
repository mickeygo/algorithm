package sort

// Insertion 插入排序
type Insertion struct {
	Elements []int
}

// Sort 插入排序。适用于少量数据的排序
//
// 规则：第 N 个元素与第 (N-1) 个元素比较，若大于，则向后第 (N+1) 与第 N 个元素比较；若小于，则 N 与第（N-2) 个元素比较，依次类推
// 时间复杂度：O(N2)
// 优点：从前向后的排序过程中，定点前的元素顺序都是已排列好的，因此当满足大于或等于前一条件时，即可退出内循环
//		最理想情况下，数组只需遍历一次，这种情况下此算法是最高效的，O(N)；最差情况下，需遍历 N+(N-1)+(N-2)+...+2+1，即 O(N2)
//		插入排序的性能很依赖与原始数组数据的顺序
// 缺点：在最差的情况下，复杂度依旧比较高
func (s *Insertion) Sort() []int {
	for i := 1; i < len(s.Elements); i++ {
		// 若发现当前位置元素不小于前一元素时，停止比较
		for j := i - 1; less(s.Elements, j+1, j); j-- {
			if less(s.Elements, j+1, j) {
				exch(s.Elements, j+1, j)
			}
		}
	}
	return s.Elements
}

func (s *Insertion) Sort2(arr []int) []int {
	return nil
}
