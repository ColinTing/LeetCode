package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question189 struct {
	name string
	para189
	ans189
}

type para189 struct {
	nums []int
	k    int
}

type ans189 struct {
	result []int
}

func Test_Problem189(t *testing.T) {
	qs := []question189 {

		{
			name: "success",
			para189: para189{nums: []int {1, 2, 3, 4, 5, 6, 7}, k: 3},
			ans189: ans189{result: []int {5, 6, 7, 1, 2, 3, 4}},
		},

		{
			name: "success",
			para189: para189{[]int{-1, -100, 3, 99}, 2},
			ans189: ans189{[]int{3, 99, -1, -100}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 189------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if rotate1(tt.para189.nums, tt.para189.k); !reflect.DeepEqual(tt.para189.nums, tt.ans189.result) {
				t.Errorf("rotate() = %v, but want %v", tt.para189.nums, tt.ans189.result)
			}
		})
	}
}