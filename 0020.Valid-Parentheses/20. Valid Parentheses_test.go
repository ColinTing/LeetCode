package leetcode

import (
	"fmt"
	"testing"
)

type question20 struct {
	name string
	para20
	ans20
}


type para20 struct {
	s string
}


type ans20 struct {
	result bool
}

func Test_Problem20(t *testing.T) {

	qs := []question20{

		{
			name: "success",
			para20: para20{"()[]{}"},
			ans20: ans20{true},
		},
		{
			name: "success",
			para20: para20{"(]"},
			ans20: ans20{false},
		},
		{
			name: "success",
			para20: para20{"({[]})"},
			ans20: ans20{true},
		},
		{
			name: "success",
			para20: para20{"(){[({[]})]}"},
			ans20: ans20{true},
		},
		{
			name: "success",
			para20: para20{"((([[[{{{"},
			ans20: ans20{false},
		},
		{
			name: "success",
			para20: para20{"(())]]"},
			ans20: ans20{false},
		},
		{
			name: "success",
			para20: para20{""},
			ans20: ans20{true},
		},
		{
			name: "success",
			para20: para20{"["},
			ans20: ans20{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 20------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.para20.s); got != tt.ans20.result {
				t.Errorf("isValid() = %v, but want %v", got, tt.ans20.result)
			}
		})
	}
}