package leetcode

import (
	"fmt"
	"testing"
)

type question169 struct {
	name string
	para169
	ans169
}


type para169 struct {
	nums []int
}


type ans169 struct {
	result int
}

func Test_Problem169(t *testing.T) {

	qs := []question169{

		{
			name: "success",
			para169: para169{[]int{2, 2, 1}},
			ans169: ans169{2},
		},

		{
			name: "success",
			para169: para169{[]int{3, 2, 3}},
			ans169: ans169{3},
		},

		{
			name: "success",
			para169: para169{[]int{2, 2, 1, 1, 1, 2, 2}},
			ans169: ans169{2},
		},

		{
			name: "success",
			para169: para169{[]int{-2147483648}},
			ans169: ans169{-2147483648},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 169------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.para169.nums); got != tt.ans169.result {
				t.Errorf("majorityElement() = %v, but want %v", got, tt.ans169.result)
			}
		})
	}
}