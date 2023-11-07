package solver

import (
	"fmt"

	"github.com/DamienDuv/godoku/models"
)

func nakedSubSet(g *models.Grid, subSetSize int) bool {
	performedAction := false
	performedAction = performedAction || nakedSubSetInSet(g.Rows, subSetSize, "row")
	performedAction = performedAction || nakedSubSetInSet(g.Cols, subSetSize, "col")
	performedAction = performedAction || nakedSubSetInSet(g.Boxes, subSetSize, "box")
	return performedAction
}

func nakedSubSetInSet(set [][]*models.Cell, subSetSize int, setType string) bool {
	performedAction := false
	for i := 0; i < 9; i++ {
		results := make([][]int, 0, 4)

		setCandidates := make([]uint16,9)
		for j := range setCandidates {
			setCandidates[j] = set[i][j].GetCandidates()
		}
		
		recurseFindSubSet(setCandidates, 0, subSetSize, subSetSize, 0, &results)

		if len(results) > 0 {
			for _, indexes := range results {
				removed := false
				pos := uint16(0)
				for _, index := range indexes {
					pos |= set[i][index].GetCandidates()
				}
				toRemove := models.BitsIndexes(pos)

				// convert the indexes to a base 1
				for m := range toRemove {
					toRemove[m]++
				}

				// remove the triple from possible candidates in the set
				for j := 0; j < 9; j++ {
					if isInArray(j, indexes) {
						continue
					}

					for _, v := range toRemove {
						removed = set[i][j].RemoveCandidateWithFeedback(v) || removed
					}
				}
				//fmt.Printf("found a naked %s %v in %s %d\n", sizeToString[subSetSize], toRemove, setType, i+1)
				if removed {
					performedAction = true
					fmt.Printf("found a naked %s %v in %s %d\n", sizeToString[subSetSize], toRemove, setType, i+1)
				}
			}
		}
	}

	return performedAction
}
