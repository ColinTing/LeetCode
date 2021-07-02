package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question589 struct {
	name string
	para589
	ans589
}


type para589 struct {
	root *Node
}


type ans589 struct {
	result []int
}

func Test_Problem589(t *testing.T) {

	qs := []question589{

		{
			name: "success",
			para589: para589{root: &Node{
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
			ans589: ans589{result: []int{1, 3, 5, 6, 2, 4}},
		},

	}

	fmt.Printf("------------------------Leetcode Problem 589------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder(tt.para589.root); !reflect.DeepEqual(got, tt.ans589.result) {
				t.Errorf("preorder() = %v, but want %v", got, tt.ans589.result)
			}
		})
	}
}

