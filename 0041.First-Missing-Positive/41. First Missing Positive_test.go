package leetcode

import (
	"fmt"
	"testing"
)

type question41 struct {
	name string
	para41
	ans41
}


type para41 struct {
	nums []int
}


type ans41 struct {
	result int
}

func Test_Problem41(t *testing.T) {

	qs := []question41{

		{
			name: "success",
			para41: para41{nums: []int{1, 2, 0}},
			ans41: ans41{result: 3},
		},

		{
			name: "success",
			para41: para41{nums: []int{3,4,-1,1}},
			ans41: ans41{result: 2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 42------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.para41.nums); got != tt.ans41.result {
				t.Errorf("firstMissingPositive() = %v, but want %v", got, tt.ans41.result)
			}
		})
	}
}