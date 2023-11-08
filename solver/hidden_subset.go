package solver

import (
	"fmt"

	"github.com/DamienDuv/godoku_solver/models"
)

func hiddenSubSet(g *models.Grid, subSetSize int) bool {
	performedAction := false
	performedAction = performedAction || hiddenSubSetInSet(g.Rows, g.DigitsPositions.Rows, subSetSize, "row")
	performedAction = performedAction || hiddenSubSetInSet(g.Cols, g.DigitsPositions.Cols, subSetSize, "col")
	performedAction = performedAction || hiddenSubSetInSet(g.Boxes, g.DigitsPositions.Boxes, subSetSize, "box")
	return performedAction
}

func hiddenSubSetInSet(set [][]*models.Cell, pos [][]uint16, subSetSize int, setType string) bool {
	performedAction := false
	for i := 0; i < 9; i++ {
		results := make([][]int, 0, 4)
		recurseFindSubSet(pos[i], 0, subSetSize, subSetSize, 0, &results)

		if len(results) > 0 {
			for _, toRemove := range results {
				removed := false

				indexMask := uint16(0)
				for _, digit := range toRemove {
					indexMask |= pos[i][digit]
				}
				indexes := models.BitsIndexes(indexMask)

				// convert the digits to a base 1
				for m := range toRemove {
					toRemove[m]++
				}

				// remove other possible candidates from the triple
				for _,ind := range indexes {
					removed = set[i][ind].RemoveCandidateExcept(toRemove...) || removed
				}
				//fmt.Printf("found a naked %s %v in %s %d\n", sizeToString[subSetSize], toRemove, setType, i+1)
				if removed {
					performedAction = true
					fmt.Printf("found a hidden %s %v in %s %d\n", sizeToString[subSetSize], toRemove, setType, i+1)
				}
			}
		}
	}

	return performedAction
}
