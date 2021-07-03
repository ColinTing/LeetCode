package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question429 struct {
	name string
	para429
	ans429
}


type para429 struct {
	root *Node
}


type ans429 struct {
	result [][]int
}

func Test_Problem589(t *testing.T) {

	qs := []question429{

		{
			name: "success",
			para429: para429{root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val:      3,
						Children: []*Node{
							{
								Val:      5,
								Children: nil,
							},
							{
								Val:      6,
								Children: nil,
							},
						},
					},
					{
						Val:      2,
						Children: nil,
					},
					{
						Val:      4,
						Children: nil,
					},
				}}},
			ans429: ans429{result: [][]int{{1},{3,2,4},{5,6}}},
		},

	}

	fmt.Printf("------------------------Leetcode Problem 589------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.para429.root); !reflect.DeepEqual(got, tt.ans429.result) {
				t.Errorf("levelOrder() = %v, but want %v", got, tt.ans429.result)
			}
		})
	}
}

