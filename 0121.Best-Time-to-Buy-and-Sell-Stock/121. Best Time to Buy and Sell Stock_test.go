package leetcode

import (
	"fmt"
	"testing"
)

type question121 struct {
	name string
	para121
	ans121
}


type para121 struct {
	prices []int
}


type ans121 struct {
	result int
}

func Test_Problem121(t *testing.T) {

	qs := []question121{

		{
			name: "success",
			para121: para121{[]int{}},
			ans121: ans121{0},
		},

		{
			name: "success",
			para121: para121{[]int{7, 1, 5, 3, 6, 4}},
			ans121: ans121{5},
		},

		{
			name: "success",
			para121: para121{[]int{7, 6, 4, 3, 1}},
			ans121: ans121{0},
		},

		{
			name: "success",
			para121: para121{[]int{1, 3, 2, 8, 4, 9}},
			ans121: ans121{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 121------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.para121.prices); got != tt.ans121.result {
				t.Errorf("maxProfit() = %v, but want %v", got, tt.ans121.result)
			}
		})
	}
}