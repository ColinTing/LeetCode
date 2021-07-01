package leetcode

import (
	"fmt"
	"testing"
)

type question242 struct {
	name string
	para242
	ans242
}


type para242 struct {
	s string
	t string
}

// ans 是答案
// one 代表第一个答案
type ans242 struct {
	result bool
}

func Test_Problem242(t *testing.T) {

	qs := []question242{

		{
			name: "success",
			para242: para242{s: "", t: ""},
			ans242: ans242{result: true},
		},
		{
			name: "success",
			para242: para242{s: "", t: "1"},
			ans242: ans242{result: false},
		},

		{
			name: "success",
			para242: para242{s: "anagram", t: "nagaram"},
			ans242: ans242{result: true},
		},

		{
			name: "success",
			para242: para242{s: "rat", t: "car"},
			ans242: ans242{result: false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 242------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.para242.s, tt.para242.t); got != tt.ans242.result {
				t.Errorf("isAnagram() = %v, but want %v", got, tt.ans242.result)
			}
		})
	}
}