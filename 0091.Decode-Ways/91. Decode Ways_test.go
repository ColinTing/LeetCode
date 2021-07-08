package leetcode

import (
	"fmt"
	"testing"
)

type question91 struct {
	name string
	para91
	ans91
}


type para91 struct {
	s string
}


type ans91 struct {
	result int
}

func Test_Problem91(t *testing.T) {

	qs := []question91{

		{
			name: "success",
			para91: para91{"12"},
			ans91: ans91{2},
		},

		{
			name: "success",
			para91: para91{"226"},
			ans91: ans91{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 91------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.para91.s); got != tt.ans91.result {
				t.Errorf("numDecodings() = %v, but want %v", got, tt.ans91.result)
			}
		})
	}
}