package solver

import (
	"fmt"
	"math/bits"

	"github.com/DamienDuv/godoku/models"
)

func nakedPair(g *models.Grid) bool {
	performedAction := false
	performedAction = performedAction || nakedPairInSet(g.Rows, "row")
	performedAction = performedAction || nakedPairInSet(g.Cols, "col")
	performedAction = performedAction || nakedPairInSet(g.Boxes, "box")
	return performedAction
}

func nakedPairInSet(set [][]*models.Cell, setType string) bool {
	performedAction := false
	for i := 0; i < 9; i++ {
		for j := 0; j < 8; j++ { // No need to check the last one as we're looking for pairs
			if bits.OnesCount16(set[i][j].GetCandidates()) != 2 {
				continue
			}

			// Only two candidates for element j, trying to find a k with same candidates
			// this will make a pair
			for k := j + 1; k < 9; k++ {
				if set[i][j].GetCandidates() != set[i][k].GetCandidates() {
					continue
				}

				// We found a pair in j/k, now we need to clear all the similar candidates in the set
				// apart from the pair
				indexes := models.BitsIndexes(set[i][j].GetCandidates())

				removed := false
				for _, ind := range indexes {
					for m := 0; m < 9; m++ {
						if m == j || m == k {
							continue
						}
						removed = set[i][m].RemoveCandidateWithFeedback(ind+1) || removed
					}
				}

				if removed {
					performedAction = true
					switch setType {
					case "row":
						fmt.Printf("found a %d/%d naked pair on row %d at col %d/%d\n", indexes[0]+1, indexes[1]+1, i+1, j+1, k+1)

					case "col":
						fmt.Printf("found a %d/%d naked pair on col %d at row %d/%d\n", indexes[0]+1, indexes[1]+1, i+1, j+1, k+1)

					case "box":
						fmt.Printf("found a %d/%d naked pair on box %d at pos %d/%d\n", indexes[0]+1, indexes[1]+1, i+1, j+1, k+1)

					}
				}

			}
		}
	}

	return performedAction
}
