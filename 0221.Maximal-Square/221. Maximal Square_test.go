package leetcode

import (
	"fmt"
	"testing"
)

type question200 struct {
	name string
	para221
	ans221
}


type para221 struct {
	matrix [][]byte
}


type ans221 struct {
	result int
}

func Test_Problem200(t *testing.T) {

	qs := []question200{

		{
			name: "success",
			para221: para221{[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			}},
			ans221: ans221{4},
		},

		{
			name: "success",
			para221: para221{[][]byte{
				{'0', '1'},
				{'1', '0'},
			}},
			ans221: ans221{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 200------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.para221.matrix); got != tt.ans221.result {
				t.Errorf("maximalSquare() = %v, but want %v", got, tt.ans221.result)
			}
		})
	}
}