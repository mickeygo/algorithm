package list

import "testing"

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