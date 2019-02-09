package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort_SortUpDown(t *testing.T) {
	arr := []int{3, 5, 7, 2, 1, 3, 4, 5, 6, 7, 8, 9, 5, 6, 3, 1}
	s := MergeSort{}
	fmt.Print(s.SortUpDown(arr))
	t.Fail()
}
