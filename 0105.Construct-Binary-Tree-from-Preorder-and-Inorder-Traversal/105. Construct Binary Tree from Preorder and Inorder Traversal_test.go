package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ColinTing/LeetCode/structures"
)

type question105 struct {
	name string
	para105
	ans105
}


type para105 struct {
	preorder []int
	inorder  []int
}


type ans105 struct {
	result []int
}

func Test_Problem105(t *testing.T) {

	qs := []question105{

		{
			name: "success",
			para105: para105{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			ans105: ans105{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 105------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := structures.Tree2ints(buildTree(tt.para105.preorder, tt.para105.inorder)); !reflect.DeepEqual(got, tt.ans105.result) {
				t.Errorf("buildTree() = %v, but want %v", got, tt.ans105.result)
			}
		})
	}
}