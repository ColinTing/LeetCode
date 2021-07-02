package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question94 struct {
	name string
	para94
	ans94
}


type para94 struct {
	root *TreeNode
}


type ans94 struct {
	result []int
}

func Test_Problem94(t *testing.T) {

	qs := []question94{

		{
			name: "success",
			para94: para94{root: nil},
			ans94: ans94{result: []int{}},
		},

		{
			name: "success",
			para94: para94{root: &TreeNode{
							Val: 1,
							Left: nil,
							Right: nil,
			}},
			ans94: ans94{result: []int{1}},
		},

		{
			name: "success",
			para94: para94{root: &TreeNode{
				Val: 1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode {
						Val: 3,
						Left: nil,
						Right: nil,
					},
					Right: nil,
				},
			}},
			ans94: ans94{result: []int{1, 3, 2}},
		},

		{
			name: "success",
			para94: para94{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{
					Val:   3,
					Left:  nil,
					Right: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
				},
				Right: nil,
			}},
			ans94: ans94{result: []int{3, 1, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 94------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.para94.root); !reflect.DeepEqual(got, tt.ans94.result) {
				t.Errorf("inorderTraversal() = %v, but want %v", got, tt.ans94.result)
			}
		})
	}
}