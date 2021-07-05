package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question47 struct {
	name string
	para47
	ans47
}


type para47 struct {
	nums []int
}


type ans47 struct {
	result [][]int
}

func Test_Problem47(t *testing.T) {

	qs := []question47{

		{
			name: "success",
			para47: para47{nums: []int{1, 1, 2}},
			ans47: ans47{result: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		},

		{
			name: "success",
			para47: para47{nums: []int{1, 2, 2}},
			ans47: ans47{result: [][]int{{1, 2, 2}, {2, 1, 2}, {2, 2, 1}}},
		},

		{
			name: "success",
			para47: para47{nums: []int{2, 2, 2}},
			ans47: ans47{result: [][]int{{2, 2, 2}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 47------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.para47.nums); !reflect.DeepEqual(got, tt.ans47.result) {
				t.Errorf("permuteUnique() = %v, but want %v", got, tt.ans47.result)
			}
		})
	}
}