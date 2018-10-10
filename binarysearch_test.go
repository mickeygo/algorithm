package algorithm

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arrs []int
		key  int
	}
	tests := []struct {
		name      string
		args      args
		wantInt   int
		wantCount int
	}{
		// TODO: Add test cases.
		{
			"BinarySearch",
			args{
				[]int{1, 2, 4, 5, 6, 8, 9, 13, 16, 18, 20},
				2,
			},
			1,
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInt, gotCount := BinarySearch(tt.args.arrs, tt.args.key)
			if gotInt != tt.wantInt {
				t.Errorf("BinarySearch() gotInt = %v, want %v", gotInt, tt.wantInt)
			}
			if gotCount != tt.wantCount {
				t.Errorf("BinarySearch() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
