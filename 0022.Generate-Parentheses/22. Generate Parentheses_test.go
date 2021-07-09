package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question22 struct {
	name string
	para22
	ans22
}

type para22 struct {
	n int
}

type ans22 struct {
	result []string
}

func Test_Problem22(t *testing.T) {

	qs := []question22{

		{
			name: "success",
			para22: para22{3},
			ans22: ans22{[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 22------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.para22.n); !reflect.DeepEqual(got, tt.ans22.result) {
				t.Errorf("generateParenthesis() = %v, but want %v", got, tt.ans22.result)
			}
		})
	}
}