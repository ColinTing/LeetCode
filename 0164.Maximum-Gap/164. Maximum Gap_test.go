package leetcode

import (
	"fmt"
	"testing"
)

type question164 struct {
	name string
	para164
	ans164
}


type para164 struct {
	nums []int
}


type ans164 struct {
	result int
}

func Test_Problem164(t *testing.T) {

	qs := []question164{

		{
			name: "success",
			para164: para164{[]int{3, 6, 9, 1}},
			ans164: ans164{3},
		},
		{
			name: "success",
			para164: para164{[]int{1}},
			ans164: ans164{0},
		},

		{
			name: "success",
			para164: para164{[]int{}},
			ans164: ans164{0},
		},

		{
			name: "success",
			para164: para164{[]int{2, 10}},
			ans164: ans164{8},
		},

		{
			name: "success",
			para164: para164{[]int{2, 435, 214, 64321, 643, 7234, 7, 436523, 7856, 8}},
			ans164: ans164{372202},
		},

		{
			name: "success",
			para164: para164{[]int{1, 10000000}},
			ans164: ans164{9999999},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 164------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGap(tt.para164.nums); got != tt.ans164.result {
				t.Errorf("maximumGap() = %v, but want %v", tt.para164.nums, tt.ans164.result)
			}
		})
	}
}