package leetcode

import (
	"fmt"
	"testing"
)

type question55 struct {
	name string
	para55
	ans55
}


type para55 struct {
	nums []int
}


type ans55 struct {
	result bool
}

func Test_Problem55(t *testing.T) {

	qs := []question55{

		{
			name: "success",
			para55: para55{[]int{2, 3, 1, 1, 4}},
			ans55: ans55{true},
		},
		{
			name: "success",
			para55: para55{[]int{3, 2, 1, 0, 4}},
			ans55: ans55{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 55------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.para55.nums); got != tt.ans55.result {
				t.Errorf("canJump() = %v, but want %v", got, tt.ans55.result)
			}
		})
	}
}