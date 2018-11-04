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
