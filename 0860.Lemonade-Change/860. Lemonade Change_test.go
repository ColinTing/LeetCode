package leetcode

import (
	"fmt"
	"testing"
)

type question860 struct {
	name string
	para860
	ans860
}

type para860 struct {
	bills   []int
}

type ans860 struct {
	result bool
}

func Test_Problem1(t *testing.T) {

	qs := []question860{
		{
			name:  "success",
			para860: para860{bills: []int{5, 5, 5, 10, 20}},
			ans860:  ans860{result: true},
		},

		{
			name:  "success",
			para860: para860{bills: []int{5, 5, 10}},
			ans860:  ans860{result: true},
		},

		{
			name:  "success",
			para860: para860{bills: []int{10, 10}},
			ans860:  ans860{result: false},
		},

		{
			name:  "success",
			para860: para860{bills: []int{5, 5, 10, 10, 20}},
			ans860:  ans860{result: false},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.para860.bills); got != tt.ans860.result {
				t.Errorf("lemonadeChange() = %v, but want %v", got, tt.ans860.result)
			}
		})
	}

}
