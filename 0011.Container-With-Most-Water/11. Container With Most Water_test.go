package leetcode

import (
	"fmt"
	"testing"
)

type question11 struct {
	name string
	para11
	ans11
}


type para11 struct {
	height []int
}


type ans11 struct {
	result int
}

func Test_Problem11(t *testing.T) {

	qs := []question11{

		{
			name: "success",
			para11: para11{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			ans11: ans11{result: 49},
		},

		{
			name: "success",
			para11: para11{height: []int{1, 1}},
			ans11: ans11{result: 1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 11------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.para11.height); got != tt.ans11.result {
				t.Errorf("maxArea() = %v, but want %v", got, tt.ans11.result)
			}
		})
	}
}