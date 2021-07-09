package leetcode

import (
	"fmt"
	"testing"
)

type question32 struct {
	name string
	para32
	ans32
}


type para32 struct {
	s string
}


type ans32 struct {
	result int
}

func Test_Problem32(t *testing.T) {

	qs := []question32{

		{
			name: "success",
			para32: para32{"(()"},
			ans32: ans32{2},
		},

		{
			name: "success",
			para32: para32{")()())"},
			ans32: ans32{4},
		},

		{
			name: "success",
			para32: para32{"()(())"},
			ans32: ans32{6},
		},

		{
			name: "success",
			para32: para32{"()(())))"},
			ans32: ans32{6},
		},

		{
			name: "success",
			para32: para32{"()(()"},
			ans32: ans32{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 32------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.para32.s); got != tt.ans32.result {
				t.Errorf("longestValidParentheses() = %v, but want %v", got, tt.ans32.result)
			}
		})
	}
}