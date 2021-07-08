package leetcode

import (
	"fmt"
	"testing"
)

type question122 struct {
	name string
	para122
	ans122
}


type para122 struct {
	prices []int
}


type ans122 struct {
	result int
}

func Test_Problem122(t *testing.T) {

	qs := []question122{

		{
			name: "success",
			para122: para122{[]int{}},
			ans122: ans122{0},
		},

		{
			name: "success",
			para122: para122{[]int{7, 1, 5, 3, 6, 4}},
			ans122: ans122{7},
		},

		{
			name: "success",
			para122: para122{[]int{7, 6, 4, 3, 1}},
			ans122: ans122{0},
		},

		{
			name: "success",
			para122: para122{[]int{1, 2, 3, 4, 5}},
			ans122: ans122{4},
		},

		{
			name: "success",
			para122: para122{[]int{1, 2, 10, 11, 12, 15}},
			ans122: ans122{14},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 122------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.para122.prices); got != tt.ans122.result {
				t.Errorf("maxProfit() = %v, but want %v", got, tt.ans122.result)
			}
		})
	}
}