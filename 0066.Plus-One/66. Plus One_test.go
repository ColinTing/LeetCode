package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question66 struct {
	name string
	para66
	ans66
}


type para66 struct {
	digits []int
}


type ans66 struct {
	result []int
}

func Test_Problem66(t *testing.T) {

	qs := []question66{

		{
			name: "success",
			para66: para66{digits: []int{1, 2, 3}},
			ans66: ans66{result: []int{1, 2, 4}},
		},

		{
			name: "success",
			para66: para66{digits: []int{4, 3, 2, 1}},
			ans66: ans66{result: []int{4, 3, 2, 2}},
		},

		{
			name: "success",
			para66: para66{digits: []int{9, 9}},
			ans66: ans66{result: []int{1, 0, 0}},
		},

		{	name: "success",
			para66: para66{digits: []int{0}},
			ans66: ans66{result: []int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 66------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.para66.digits); !reflect.DeepEqual(got, tt.ans66.result) {
				t.Errorf("plusOne() = %v, but want %v", got, tt.ans66.result)
			}
		})
	}
}