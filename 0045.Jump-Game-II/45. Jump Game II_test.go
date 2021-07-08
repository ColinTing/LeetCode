package leetcode

import (
	"fmt"
	"testing"
)

type question45 struct {
	name string
	para45
	ans45
}


type para45 struct {
	nums []int
}


type ans45 struct {
	result int
}

func Test_Problem45(t *testing.T) {

	qs := []question45{

		{
			name: "success",
			para45: para45{[]int{2, 3, 1, 1, 4}},
			ans45: ans45{2},
		},

		{
			name: "success",
			para45: para45{[]int{2, 3, 0, 1, 4}},
			ans45: ans45{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 45------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.para45.nums); got != tt.ans45.result {
				t.Errorf("jump() = %v, but want %v", got, tt.ans45.result)
			}
		})
	}
}