package leetcode

import (
	"fmt"
	"testing"
)

type question64 struct {
	name string
	para64
	ans64
}


type para64 struct {
	grid [][]int
}


type ans64 struct {
	result int
}

func Test_Problem64(t *testing.T) {

	qs := []question64{

		{
			name: "success",
			para64: para64{[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			}},
			ans64: ans64{7},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 64------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.para64.grid); got != tt.ans64.result {
				t.Errorf("minPathSum() = %v, but want %v", got, tt.ans64.result)
			}
		})
	}
}