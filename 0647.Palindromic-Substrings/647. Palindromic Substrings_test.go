package leetcode

import (
	"fmt"
	"testing"
)

type question647 struct {
	name string
	para647
	ans647
}


type para647 struct {
	s string
}


type ans647 struct {
	result int
}

func Test_Problem647(t *testing.T) {

	qs := []question647{

		{
			name: "success",
			para647: para647{"abc"},
			ans647: ans647{3},
		},

		{
			name: "success",
			para647: para647{"aaa"},
			ans647: ans647{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 647------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.para647.s); got != tt.ans647.result {
				t.Errorf("countSubstrings() = %v, but want %v", got, tt.ans647.result)
			}
		})
	}
}