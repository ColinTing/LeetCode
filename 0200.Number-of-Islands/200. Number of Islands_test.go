package leetcode

import (
	"fmt"
	"testing"
)

type question200 struct {
	name string
	para200
	ans200
}


type para200 struct {
	grid [][]byte
}


type ans200 struct {
	result int
}

func Test_Problem200(t *testing.T) {

	qs := []question200{

		{
			name: "success",
			para200: para200{[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}},
			ans200: ans200{1},
		},

		{
			name: "success",
			para200: para200{[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			}},
			ans200: ans200{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 200------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.para200.grid); got != tt.ans200.result {
				t.Errorf("numIslands() = %v, but want %v", got, tt.ans200.result)
			}
		})
	}
}