package lc

import "testing"

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
		output int
	}{}
}
