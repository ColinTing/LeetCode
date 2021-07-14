package leetcode

import (
	"fmt"
	"testing"
)

type question152 struct {
	name string
	para152
	ans152
}


type para152 struct {
	nums []int
}


type ans152 struct {
	result int
}

func Test_Problem152(t *testing.T) {

	qs := []question152{

		{
			name: "success",
			para152: para152{[]int{-2}},
			ans152: ans152{-2},
		},

		{
			name: "success",
			para152: para152{[]int{3, -1, 4}},
			ans152: ans152{4},
		},

		{
			name: "success",
			para152: para152{[]int{0}},
			ans152: ans152{0},
		},

		{
			name: "success",
			para152: para152{[]int{2, 3, -2, 4}},
			ans152: ans152{6},
		},

		{
			name: "success",
			para152: para152{[]int{-2, 0, -1}},
			ans152: ans152{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 152------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.para152.nums); got != tt.ans152.result {
				t.Errorf("maxProduct() = %v, but want %v", got, tt.ans152.result)
			}
		})
	}
}