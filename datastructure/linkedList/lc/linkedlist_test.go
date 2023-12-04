package lc

import (
	"testing"
)

func linkedListEqual(l1, l2 *ListNode) bool {
	equal := true
	for {
		if l1.Val != l2.Val {
			equal = false
			break
		}

		if l1.Next == nil || l2.Next == nil {
			if l1.Next != l2.Next {
				equal = false
			}
			break
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return equal
}

func Test_MiddleNode(t *testing.T) {
	tests := []struct {
		input  *ListNode
		output int
	}{
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			output: 3,
		},
	}

	for _, v := range tests {
		res := middleNode(v.input)
		if res.Val != v.output {
			t.Errorf("Got %+v, expected %+v", res, v.output)
		}
	}

}
func Test_DeleteDuplicate(t *testing.T) {
	tests := []struct {
		input  *ListNode
		output *ListNode
	}{
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
		{
			input:  &ListNode{},
			output: &ListNode{},
		},
	}

	for _, v := range tests {
		res := deleteDuplicates(v.input)
		if !linkedListEqual(res, v.output) {
			t.Errorf("Got %+v, expected %+v", res, v.output)
		}
	}
}

func Test_ReverseBetween(t *testing.T) {
	tests := []struct {
		input  *ListNode
		left   int
		right  int
		output *ListNode
	}{
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			left:  2,
			right: 4,
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			input: &ListNode{
				Val:  5,
				Next: nil,
			},
			left:  1,
			right: 1,
			output: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			left:  2,
			right: 3,
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	for _, v := range tests {
		res := reverseBetween(v.input, v.left, v.right)
		if !linkedListEqual(res, v.output) {
			t.Errorf("Got %+v, expected %+v", res, v.output)
		}
	}
}
