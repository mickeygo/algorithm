package sort

import (
	"fmt"
	"testing"
)

func TestShell_Sort(t *testing.T) {
	shell := &Shell{}
	arr := shell.Sort([]int{1, 3, 5, 7, 9, 8, 0, 9, 45, 56, 43, 12, 34, 56, 54, 12, 34, 65, 90})
	fmt.Print(arr)

	t.Fail()
}

func TestShell_Sort2(t *testing.T) {
	shell := &Shell{}
	arr := shell.Sort2([]int{1, 3, 5, 7, 9, 8, 0, 9, 45, 56, 43, 12, 34, 56, 54, 12, 34, 65, 90})
	fmt.Print(arr)

	t.Fail()
}
