package solver

import (
	"math/bits"
)

// recurseFindSubSet recursively try to find subSets (pair/triplet/quadruplet...) in a set (row/column/box)
// the first call has to be made wit start and mask at 0 and setSize and depth at the desired subSet size
// if it finds one or more subSet it will populate the results array with arrays containing
// the indexes of the subset elements in the set.
// the return value is only use as internal purpose of the recursion
// NOTE: recurseFindSubSet WILL NOT YIELD PROPER RESULTS IF NAKED SINGLE / HIDDEN SINGLES ARE LEFT IN THE SET
func recurseFindSubSet(set []uint16, start int, setSize int, depth int, mask uint16, results *[][]int) *[]int {
	if depth == 0 {
		if bits.OnesCount16(mask) == setSize {
			*results = append(*results, make([]int, 0, setSize))
			return &((*results)[len(*results)-1])
		}
		return nil
	}

	for i := start; i <= 9-depth; i++ {
		if set[i] == 0 {
			continue
		}

		if bits.OnesCount16(set[i]|mask) <= setSize {
			newMask := set[i] | mask
			resultSet := recurseFindSubSet(set, i+1, setSize, depth-1, newMask, results)

			if resultSet != nil {
				*resultSet = append(*resultSet, i)
				if depth != setSize {
					return resultSet
				}
			}
		}
	}

	return nil
}