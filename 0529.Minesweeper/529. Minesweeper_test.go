package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question529 struct {
	name string
	para529
	ans529
}


type para529 struct {
	board     [][]byte
	click []int
}


type ans529 struct {
	result [][]byte
}

func Test_Problem529(t *testing.T) {

	qs := []question529{

		{
			name: "success",
			para529: para529{[][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			}, []int{3, 0}},
			ans529: ans529{[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			}},
		},

		{
			name: "success",
			para529: para529{[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			}, []int{1, 2}},
			ans529: ans529{[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'X', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 529------------------------\n")

	for _, tt := range qs {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateBoard(tt.para529.board, tt.para529.click); !reflect.DeepEqual(got, tt.ans529.result) {
				t.Errorf("updateBoard() = %v, but want %v", got, tt.ans529.result)
			}
		})
	}
}