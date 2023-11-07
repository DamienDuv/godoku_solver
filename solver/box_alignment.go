package solver

import (
	"fmt"

	"github.com/DamienDuv/godoku/models"
)

// Finds an alignment of the same candidate withing a box and clears the corresponding
// row or column of this candidate (outside of the box containing the alignment)
func alignmentInBox(g *models.Grid) bool {
	performedAction := false

	alignments := []uint16{
		0b000000111,
		0b000111000,
		0b111000000,
		0b001001001,
		0b010010010,
		0b100100100,
	}


	for i := 0; i < 9; i++ {
		for val := 1; val <= 9; val++ {
			posInBox := g.DigitsPositions.Boxes[i][val-1]
			if posInBox == 0 { // digit already in box
				continue
			}

			for k, a := range alignments {
				if posInBox&a != posInBox { // values are not on the current alignment mask
					continue
				}

				if k < 3 { // is on a row alignment
					row := (i/3)*3 + k
					removed := false
					for l := 0; l < 9; l++ {
						if l < (i%3)*3 || l >= (i%3)*3+3 {
							res := g.Rows[row][l].RemoveCandidateWithFeedback(val)
							removed = removed || res
						}
					}

					if removed {
						fmt.Printf("alignment of %d in box %d, row %d\n", val, i+1, row+1)
						performedAction = true
					}
				} else { // is on a column alignment
					col := (i%3)*3 + k - 3
					removed := false
					for l := 0; l < 9; l++ {
						if l < (i/3)*3 || l >= (i/3)*3+3 {
							res := g.Cols[col][l].RemoveCandidateWithFeedback(val)
							removed = removed || res
						}
					}

					if removed {
						fmt.Printf("alignment of %d in box %d, col %d\n", val, i+1, col+1)
						performedAction = true
					}

				}
				break

			}
		}
	}

	return performedAction
}
