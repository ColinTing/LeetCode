package leetcode

import (
	"fmt"
	"testing"
)

type question153 struct {
	name string
	para153
	ans153
}


type para153 struct {
	nums []int
}


type ans153 struct {
	result int
}

func Test_Problem153(t *testing.T) {

	qs := []question153{

		{
			name: "success",
			para153: para153{[]int{5, 1, 2, 3, 4}},
			ans153: ans153{1},
		},

		{
			name: "success",
			para153: para153{[]int{1}},
			ans153: ans153{1},
		},

		{
			name: "success",
			para153: para153{[]int{1, 2}},
			ans153: ans153{1},
		},

		{
			name: "success",
			para153: para153{[]int{2, 1}},
			ans153: ans153{1},
		},

		{
			name: "success",
			para153: para153{[]int{2, 3, 1}},
			ans153: ans153{1},
		},

		{
			name: "success",
			para153: para153{[]int{1, 2, 3}},
			ans153: ans153{1},
		},

		{
			name: "success",
			para153: para153{[]int{3, 4, 5, 1, 2}},
			ans153: ans153{1},
		},

		{
			name: "success",
			para153: para153{[]int{4, 5, 6, 7, 0, 1, 2}},
			ans153: ans153{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 153------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.para153.nums); got != tt.ans153.result {
				t.Errorf("findMin() = %v, but want %v", got, tt.ans153.result)
			}
		})
	}
}