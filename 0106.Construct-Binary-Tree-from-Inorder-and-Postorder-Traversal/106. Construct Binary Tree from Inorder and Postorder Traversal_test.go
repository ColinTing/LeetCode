package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ColinTing/LeetCode/structures"
)

type question106 struct {
	name string
	para106
	ans106
}


type para106 struct {
	inorder   []int
	postorder []int
}


type ans106 struct {
	result []int
}

func Test_Problem106(t *testing.T) {

	qs := []question106{

		{
			name: "success",
			para106: para106{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}},
			ans106: ans106{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 106------------------------\n")




	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := structures.Tree2ints(buildTree(tt.para106.inorder, tt.para106.postorder)); !reflect.DeepEqual(got, tt.ans106.result) {
				t.Errorf("buildTree() = %v, but want %v", got, tt.ans106.result)
			}
		})
	}
}