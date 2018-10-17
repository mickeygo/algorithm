package sort

import (
	"reflect"
	"testing"
)

func TestInsertion_Sort(t *testing.T) {
	tests := []struct {
		name string
		s    *Insertion
		want []int
	}{
		// TODO: Add test cases.
		{
			"Insertion_1",
			&Insertion{
				Elements: []int{3, 4, 5, 6, 6, 7, 3, 8},
			},
			[]int{3, 3, 4, 5, 6, 6, 7, 8},
		},
		{
			"Insertion_2",
			&Insertion{
				Elements: []int{8},
			},
			[]int{8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insertion.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertion_Sort2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		s    *Insertion
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Sort2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insertion.Sort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
