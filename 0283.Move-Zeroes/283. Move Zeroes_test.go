package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question283 struct {
	name string
	para283
	ans283
}


type para283 struct {
	nums []int
}


type ans283 struct {
	result []int
}

func Test_Problem283(t *testing.T) {

	qs := []question283{

		{
			name: "success",
			para283: para283{[]int{1, 0, 1}},
			ans283: ans283{[]int{1, 1, 0}},
		},

		{
			name: "success",
			para283: para283{[]int{0, 1, 0, 3, 0, 12}},
			ans283: ans283{[]int{1, 3, 12, 0, 0, 0}},
		},

		{
			name: "success",
			para283: para283{[]int{0, 1, 0, 3, 0, 0, 0, 0, 1, 12}},
			ans283: ans283{[]int{1, 3, 1, 12, 0, 0, 0, 0, 0, 0}},
		},

		{
			name: "success",
			para283: para283{[]int{0, 0, 0, 0, 0, 0, 0, 0, 12, 1}},
			ans283: ans283{[]int{12, 1, 0, 0, 0, 0, 0, 0, 0, 0}},
		},

		{
			name: "success",
			para283: para283{[]int{0, 0, 0, 0, 0}},
			ans283: ans283{[]int{0, 0, 0, 0, 0}},
		},

		{
			name: "success",
			para283: para283{[]int{1}},
			ans283: ans283{[]int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 283------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if moveZeroes1(tt.para283.nums); !reflect.DeepEqual(tt.para283.nums, tt.ans283.result) {
				t.Errorf("moveZeroes() = %v, but want %v", tt.para283.nums, tt.ans283.result)
			}
		})
	}
}