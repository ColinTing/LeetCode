package leetcode

import (
	"fmt"
	"testing"
)

type question49 struct {
	name string
	para49
	ans49
}


type para49 struct {
	strs []string
}


type ans49 struct {
	result [][]string
}

func Test_Problem49(t *testing.T) {

	qs := []question49{

		{
			name:  "success",
			para49: para49{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			ans49: ans49{result: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 49------------------------\n")

	for _, q := range qs {
		a, p := q.ans49, q.para49
		fmt.Printf("groupAnagrams() = %v  maybe true, \n" +
			"if result unorder = %v",  groupAnagrams(p.strs), a.result)
	}
	fmt.Printf("\n\n\n")
}