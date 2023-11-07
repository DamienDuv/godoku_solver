package solver

import (
	"fmt"
	"math/bits"

	"github.com/DamienDuv/godoku/models"
)

func hiddenPair(g *models.Grid) bool {
	performedAction := false
	performedAction = performedAction || hiddenPairInSet(g.Rows, g.DigitsPositions.Rows, "row")
	performedAction = performedAction || hiddenPairInSet(g.Cols, g.DigitsPositions.Cols, "col")
	performedAction = performedAction || hiddenPairInSet(g.Boxes, g.DigitsPositions.Boxes, "box")
	return performedAction
}

func hiddenPairInSet(gridSet [][]*models.Cell, posSet [][]uint16, setType string) bool {
	performedAction := false
	for i := 0; i < 9; i++ {
		// look for two digits v and d having the same mask on the row
		// and have only two occurences. if we found one, it is a pair
		for v := 1; v <= 9; v++ {
			for d := v + 1; d <= 9; d++ {
				if posSet[v-1][i] != posSet[d-1][i] {
					continue
				}

				if bits.OnesCount16(posSet[v-1][i]) != 2 {
					continue
				}
				if bits.OnesCount16(posSet[d-1][i]) != 2 {
					continue
				}

				// we have a pair (no pun intended)
				indexes := models.BitsIndexes(posSet[v-1][i])

				removed := false
				for _, ind := range indexes {
					res := gridSet[i][ind].RemoveCandidateExcept(v, d)
					removed = removed || res
				}

				if removed {
					performedAction = true
					switch setType {
					case "row":
						fmt.Printf("found a %d/%d hidden pair on row %d at col %d/%d\n", v, d, i+1, indexes[0]+1, indexes[1]+1)

					case "col":
						fmt.Printf("found a %d/%d hidden pair on col %d at row %d/%d\n", v, d, i+1, indexes[0]+1, indexes[1]+1)

					case "box":
						fmt.Printf("found a %d/%d hidden pair on box %d at pos %d/%d\n", v, d, i+1, indexes[0]+1, indexes[1]+1)

					}

				}

			}
		}
	}
	return performedAction
}
