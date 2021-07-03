package leetcode

import (
	"fmt"
	"testing"
)

type question264 struct {
	name string
	para264
	ans264
}


type para264 struct {
	n int
}


type ans264 struct {
	result int
}

func Test_Problem264(t *testing.T) {

	qs := []question264{

		{
			name: "success",
			para264: para264{n: 10},
			ans264: ans264{result: 12},
		},

		{
			name: "success",
			para264: para264{n: 1},
			ans264: ans264{result: 1},
		},

		{
			name: "success",
			para264: para264{n: 6},
			ans264: ans264{result: 6},
		},

		{
			name: "success",
			para264: para264{n: 8},
			ans264: ans264{result: 9},
		},

		{
			name: "success",
			para264: para264{n: 14},
			ans264: ans264{result: 20},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 264------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.para264.n); got != tt.ans264.result {
				t.Errorf("nthUglyNumber() = %v, but want %v", got, tt.ans264.result)
			}
		})
	}
}