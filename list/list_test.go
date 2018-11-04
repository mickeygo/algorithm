package list

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	node := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: nil,
			},
		},
	}

	fmt.Print(Search(node, 2))
	t.Fail()
}

func TestDelete(t *testing.T) {
	node := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: nil,
			},
		},
	}

	Print(Delete(node, 5))
	t.Fail()
}

func TestReverse(t *testing.T) {
	node := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: nil,
			},
		},
	}

	tail := Reverse(node)

	Print(tail)

	t.Fail()
}

func TestUnionSorted(t *testing.T) {
	node1 := &NodeRune{
		Data: 1,
		Next: &NodeRune{
			Data: 3,
			Next: &NodeRune{
				Data: 5,
				Next: &NodeRune{
					Data: 7,
					Next: nil,
				},
			},
		},
	}

	node2 := &NodeRune{
		Data: 2,
		Next: &NodeRune{
			Data: 4,
			Next: &NodeRune{
				Data: 6,
				Next: nil,
			},
		},
	}

	node := UnionSorted(node1, node2)
	for node != nil {
		fmt.Printf("%d ... \n", node.Data)
		node = node.Next
	}

	t.Fail()
}
