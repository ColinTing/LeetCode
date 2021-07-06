package leetcode

import (
	"fmt"
	"testing"
)

type question127 struct {
	name string
	para127
	ans127
}


type para127 struct {
	beginWord string
	endWord string
	wordList []string
}


type ans127 struct {
	result int
}

func Test_Problem127(t *testing.T) {

	qs := []question127{
		{
			name: "success",
			para127: para127{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
			ans127: ans127{5},
		},

		{
			name: "success",
			para127: para127{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}},
			ans127: ans127{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 127------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.para127.beginWord, tt.para127.endWord, tt.para127.wordList); got != tt.ans127.result {
				t.Errorf("ladderLength() = %v, but want %v", got, tt.ans127.result)
			}
		})
	}
}