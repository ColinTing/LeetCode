package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ColinTing/LeetCode/structures"
)

type question21 struct {
	name string
	para21
	ans21
}

type para21 struct {
	l1 []int
	l2 []int
}

type ans21 struct {
	result []int
}

func Test_Problem21(t *testing.T) {

	qs := []question21{

		{
			name:   "success",
			para21: para21{l1: []int{}, l2: []int{}},
			ans21:  ans21{result: []int{}},
		},

		{
			name:   "success",
			para21: para21{l1: []int{1}, l2: []int{1}},
			ans21:  ans21{result: []int{1, 1}},
		},

		{
			name:   "success",
			para21: para21{l1: []int{1, 2, 3, 4}, l2: []int{1, 2, 3, 4}},
			ans21:  ans21{result: []int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		{
			name:   "success",
			para21: para21{l1: []int{1, 2, 3, 4, 5}, l2: []int{1, 2, 3, 4, 5}},
			ans21:  ans21{result: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		{
			name:   "success",
			para21: para21{l1: []int{1}, l2: []int{9, 9, 9, 9, 9}},
			ans21:  ans21{result: []int{1, 9, 9, 9, 9, 9}},
		},

		{
			name:   "success",
			para21: para21{l1: []int{9, 9, 9, 9, 9}, l2: []int{1}},
			ans21:  ans21{result: []int{1, 9, 9, 9, 9, 9}},
		},

		{
			name:   "success",
			para21: para21{l1: []int{2, 3, 4}, l2: []int{4, 5, 6}},
			ans21:  ans21{result: []int{2, 3, 4, 4, 5, 6}},
		},

		{
			name:   "success",
			para21: para21{l1: []int{1, 3, 8}, l2: []int{1, 7}},
			ans21:  ans21{result: []int{1, 1, 3, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 21------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := structures.List2Ints(mergeTwoLists(structures.Ints2List(tt.para21.l1), structures.Ints2List(tt.para21.l2))); !reflect.DeepEqual(got, tt.ans21.result) {
				t.Errorf("mergeTwoLists() = %v, but want %v", got, tt.ans21.result)
			}
		})
	}
}
