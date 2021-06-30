package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question88 struct {
	name string
	para88
	ans88
}

// para 是参数
// one 代表第一个参数
type para88 struct {
	nums1 []int
	m   int
	nums2 []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans88 struct {
	result []int
}

func Test_Problem88(t *testing.T) {

	qs := []question88{

		{
			name: "success",
			para88: para88{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1},
			ans88: ans88{result: []int{1}},
		},

		{
			name: "success",
			para88: para88{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3},
			ans88: ans88{result: []int{1, 2, 2, 3, 5, 6}},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 88------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if merge(tt.para88.nums1, tt.para88.m, tt.para88.nums2, tt.para88.n); !reflect.DeepEqual(tt.para88.nums1, tt.ans88.result) {
				t.Errorf("merge() = %v, but want %v", tt.para88.nums1, tt.ans88.result)
			}
		})
	}
}