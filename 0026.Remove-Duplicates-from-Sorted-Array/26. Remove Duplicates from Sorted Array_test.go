package leetcode

import (
	"fmt"
	"testing"
)

type question26 struct {
	name string
	para26
	ans26
}

type para26 struct {
	nums []int
}

type ans26 struct {
	result int
}

func Test_Problem26(t *testing.T) {

	qs := []question26{

		{
			name:   "success",
			para26: para26{nums: []int{1, 1, 2}},
			ans26:  ans26{result: 2},
		},

		{
			name:   "success",
			para26: para26{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 4, 4}},
			ans26:  ans26{result: 5},
		},

		{
			name:   "success",
			para26: para26{nums: []int{0, 0, 0, 0, 0}},
			ans26:  ans26{result: 1},
		},

		{
			name:   "success",
			para26: para26{nums: []int{1}},
			ans26:  ans26{result: 1},
		},
	}

	fmt.Print("------------------------Leetcode Problem 26------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates1(tt.para26.nums); got != tt.ans26.result {
				t.Errorf("removeDuplicates() = %v, but want %v", got, tt.ans26.result)
			}
		})
	}
}
