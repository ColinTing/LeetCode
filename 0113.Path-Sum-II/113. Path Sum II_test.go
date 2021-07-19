package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ColinTing/LeetCode/structures"
)

type question113 struct {
	name string
	para113
	ans113
}


type para113 struct {
	root []int
	targetSum int
}


type ans113 struct {
	result [][]int
}

func Test_Problem113(t *testing.T) {

	qs := []question113{
		{
			name: "success",
			para113: para113{[]int{1, 0, 1, 1, 2, 0, -1, 0, 1, -1, 0, -1, 0, 1, 0}, 2},
			ans113: ans113{[][]int{{1, 0, 1, 0}, {1, 0, 2, -1}, {1, 1, 0, 0}, {1, 1, -1, 1}}},
		},

		{
			name: "success",
			para113: para113{[]int{5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, structures.NULL, structures.NULL, 5, 1}, 22},
			ans113: ans113{[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 113------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(structures.Ints2TreeNode(tt.para113.root), tt.para113.targetSum); !reflect.DeepEqual(got, tt.ans113.result) {
				t.Errorf("pathSum() = %v, but want %v", got, tt.ans113.result)
			}
		})
	}

}