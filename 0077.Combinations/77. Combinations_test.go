package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question77 struct {
	name string
	para77
	ans77
}


type para77 struct {
	n int
	k int
}


type ans77 struct {
	result [][]int
}

func Test_Problem77(t *testing.T) {

	qs := []question77{

		{
			name: "success",
			para77: para77{n: 4, k: 2},
			ans77: ans77{result: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 77------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.para77.n, tt.para77.k); !reflect.DeepEqual(got, tt.ans77.result) {
				t.Errorf("combine() = %v, but want %v", got, tt.ans77.result)
			}
		})
	}
}