package sort

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

	gap := len(arr) / 2 // 元素间隔，初始值为数组的一半; 随便也给定初始化的分组数目
	for gap > 0 {
		for i := 0; i < gap; i++ {

		}

		gap = gap >> 1 // (gap / 2)
	}

	return nil
}
