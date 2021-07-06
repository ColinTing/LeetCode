package leetcode

import (
	"fmt"
	"testing"
)

type question455 struct {
	name string
	para455
	ans455
}


type para455 struct {
	g []int
	s []int
}


type ans455 struct {
	result int
}

func Test_Problem455(t *testing.T) {

	qs := []question455{

		{
			name: "success",
			para455: para455{[]int{1, 2, 3}, []int{1, 1}},
			ans455: ans455{1},
		},

		{
			name: "success",
			para455: para455{[]int{1, 2}, []int{1, 2, 3}},
			ans455: ans455{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 455------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(tt.para455.g, tt.para455.s); got != tt.ans455.result {
				t.Errorf("findContentChildren() = %v, but want %v", got, tt.ans455.result)
			}
		})
	}
}