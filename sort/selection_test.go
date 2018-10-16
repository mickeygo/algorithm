package sort

import (
	"reflect"
	"testing"
)

func TestSelection_Sort(t *testing.T) {
	tests := []struct {
		name string
		s    *Selection
		want []int
	}{
		// TODO: Add test cases.
		{
			"Selection_1",
			&Selection{
				Elements: []int{3, 4, 5, 6, 6, 7, 3, 8},
			},
			[]int{3, 3, 4, 5, 6, 6, 7, 8},
		},
		{
			"Selection_2",
			&Selection{
				Elements: []int{3, 4, 5, 6, 6, 7, 1, 2},
			},
			[]int{1, 2, 3, 4, 5, 6, 6, 7},
		},
		{
			"Selection_3",
			&Selection{
				Elements: []int{3},
			},
			[]int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Selection.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
