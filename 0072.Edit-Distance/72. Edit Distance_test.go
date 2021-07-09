package leetcode

import (
	"fmt"
	"testing"
)

type question74 struct {
	name string
	para72
	ans72
}


type para72 struct {
	word1 string
	word2 string
}


type ans72 struct {
	result int
}

func Test_Problem74(t *testing.T) {

	qs := []question74{

		{
			name: "success",
			para72: para72{"horse", "ros"},
			ans72: ans72{3},
		},

		{
			name: "success",
			para72: para72{"intention", "execution"},
			ans72: ans72{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 74------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.para72.word1, tt.para72.word2); got != tt.ans72.result {
				t.Errorf("searchMatrix() = %v, but want %v", got, tt.ans72.result)
			}
		})
	}
}