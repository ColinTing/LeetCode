package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question1 struct {
	name string
	para1
	ans1
}

type para1 struct {
	nums   []int
	target int
}

type ans1 struct {
	result []int
}

func Test_Problem1(t *testing.T) {

	qs := []question1{
		{
			name:  "success",
			para1: para1{nums: []int{3, 2, 4}, target: 6},
			ans1:  ans1{result: []int{1, 2}},
		},

		{
			name:  "success",
			para1: para1{nums: []int{3, 2, 4}, target: 5},
			ans1:  ans1{result: []int{0, 1}},
		},

		{
			name:  "success",
			para1: para1{nums: []int{0, 8, 7, 3, 3, 4, 2}, target: 11},
			ans1:  ans1{result: []int{1, 3}},
		},

		{
			name:  "success",
			para1: para1{nums: []int{0, 1}, target: 1},
			ans1:  ans1{result: []int{0, 1}},
		},

		{
			name:  "success",
			para1: para1{nums: []int{0, 3}, target: 5},
			ans1:  ans1{result: []int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.para1.nums, tt.para1.target); !reflect.DeepEqual(got, tt.ans1.result) {
				t.Errorf("twoSum() = %v, but want %v", got, tt.ans1.result)
			}
		})
	}

}
