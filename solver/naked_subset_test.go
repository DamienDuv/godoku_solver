package solver

import (
	"testing"
)

type FindTripleTest struct {
	candidates []uint16
	result     [][]int
}

var testSet = []FindTripleTest{
	{
		candidates: []uint16{
			0b000000000,
			0b001100000,
			0b001110000,
			0b000000000,
			0b000000000,
			0b000000000,
			0b000000000,
			0b001010000,
			0b000000000,
		},
		result: [][]int{{7, 2, 1}},
	},
	{
		candidates: []uint16{
			0b000000000,
			0b001110000,
			0b001110000,
			0b000000000,
			0b000000000,
			0b000000000,
			0b000000000,
			0b001000100,
			0b000000000,
		},
		result: [][]int{},
	},
	{
		candidates: []uint16{
			0b000000000,
			0b001110000,
			0b001110000,
			0b000000000,
			0b111111111,
			0b000000000,
			0b000000000,
			0b000000000,
			0b001010000,
		},
		result: [][]int{{8, 2, 1}},
	},
	{
		candidates: []uint16{
			0b000000000,
			0b001100000,
			0b001100000,
			0b000000000,
			0b111111111,
			0b000000000,
			0b000000000,
			0b000000000,
			0b000000000,
		},
		result: [][]int{},
	},
	{
		candidates: []uint16{
			0b000000000,
			0b001110000,
			0b001100000,
			0b010000001,
			0b111111111,
			0b000000000,
			0b110000000,
			0b001010000,
			0b100000001,
		},
		result: [][]int{{7, 2, 1}, {8, 6, 3}},
	},
}

func TestFindTriple(t *testing.T) {

	for _, test := range testSet {

		results := make([][]int, 0, 3)
		recurseFindSubSet(test.candidates, 0, 3, 3, 0, &results)

		t.Logf("expected : %v, result : %v\n", test.result, results)
		if len(test.result) != len(results) {
			t.Fail()
			return
		}

		for i, res := range test.result {
			for j, v := range res {
				if results[i][j] != v {
					t.Fail()
					return
				}
			}
		}
	}
}
