package leetcode

import (
	"fmt"
	"testing"
)

type question74 struct {
	name string
	para74
	ans74
}


type para74 struct {
	matrix [][]int
	target int
}


type ans74 struct {
	result bool
}

func Test_Problem74(t *testing.T) {

	qs := []question74{

		{
			name: "success",
			para74: para74{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3},
			ans74: ans74{true},
		},

		{
			name: "success",
			para74: para74{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13},
			ans74: ans74{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 74------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.para74.matrix, tt.para74.target); got != tt.ans74.result {
				t.Errorf("searchMatrix() = %v, but want %v", got, tt.ans74.result)
			}
		})
	}
}