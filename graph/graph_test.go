package graph

import (
	"fmt"
	"testing"
)

func TestGraph_String(t *testing.T) {
	// 	*  13 vertices, 13 edges
	//  *  0: 6 2 1 5
	//  *  1: 0
	//  *  2: 0
	//  *  3: 5 4
	//  *  4: 5 6 3
	//  *  5: 3 4 0
	//  *  6: 0 4
	//  *  7: 8
	//  *  8: 7
	//  *  9: 11 10 12
	//  *  10: 9
	//  *  11: 9 12
	//  *  12: 11 9
	args := map[int][]int{
		0:  []int{6, 2, 1, 5},
		1:  []int{0},
		2:  []int{0},
		3:  []int{5, 4},
		4:  []int{5, 6, 3},
		5:  []int{3, 4, 0},
		6:  []int{0, 4},
		7:  []int{8},
		8:  []int{7},
		9:  []int{11, 10, 12},
		10: []int{9},
		11: []int{9, 12},
		12: []int{11, 9},
	}

	graph := NewGraph(args)

	fmt.Println(graph)
}
