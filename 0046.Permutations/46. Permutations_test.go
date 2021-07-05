package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question46 struct {
	name string
	para46
	ans46
}


type para46 struct {
	nums []int
}


type ans46 struct {
	result [][]int
}

func Test_Problem46(t *testing.T) {

	qs := []question46{

		{
			name: "success",
			para46: para46{nums: []int{1, 2, 3}},
			ans46: ans46{result: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 46------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.para46.nums); !reflect.DeepEqual(got, tt.ans46.result) {
				t.Errorf("permute() = %v, but want %v", got, tt.ans46.result)
			}
		})
	}
}