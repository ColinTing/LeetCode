package leetcode

import (
	"fmt"
	"testing"
)

type question42 struct {
	name string
	para42
	ans42
}


type para42 struct {
	height []int
}


type ans42 struct {
	result int
}

func Test_Problem42(t *testing.T) {

	qs := []question42{

		{
			name: "success",
			para42: para42{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			ans42: ans42{result: 6},
		},

		{
			name: "success",
			para42: para42{height: []int{1, 2}},
			ans42: ans42{result: 0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 42------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.para42.height); got != tt.ans42.result {
				t.Errorf("trap() = %v, but want %v", got, tt.ans42.result)
			}
		})
	}
}