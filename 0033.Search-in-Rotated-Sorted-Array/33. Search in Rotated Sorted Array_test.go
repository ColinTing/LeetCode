package leetcode

import (
	"fmt"
	"testing"
)

type question33 struct {
	name string
	para33
	ans33
}


type para33 struct {
	nums   []int
	target int
}


type ans33 struct {
	result int
}

func Test_Problem33(t *testing.T) {

	qs := []question33{

		{
			name: "success",
			para33: para33{[]int{3, 1}, 1},
			ans33: ans33{1},
		},

		{
			name: "success",
			para33: para33{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			ans33: ans33{4},
		},

		{
			name: "success",
			para33: para33{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			ans33: ans33{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 33------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.para33.nums, tt.para33.target); got != tt.ans33.result {
				t.Errorf("search() = %v, but want %v", got, tt.ans33.result)
			}
		})
	}
}