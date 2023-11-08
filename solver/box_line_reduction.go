package solver


import (
	"fmt"

	"github.com/DamienDuv/godoku_solver/models"
)

// boxLineReduction finds a candidate within a row/col that is confined to a single box
// and clear all other similar candidate in that box
func boxLineReduction(g *models.Grid) bool {
	performedAction := false

	alignments := []uint16{
		0b000000111,
		0b000111000,
		0b111000000,
	}

	// Rows
	for i := 0; i < 9; i++ {
		for val := 1; val <= 9; val++ {
			posOnRow := g.DigitsPositions.Rows[i][val-1]
			if posOnRow == 0 { // digit already in row
				continue
			}

			for k, a := range alignments {
				if posOnRow&a != posOnRow { // values are not on the current alignment mask
					continue
				}

				
				removed := false
				for l := 0; l < 3; l++ {	
					for m := 0; m < 3; m++ {
						ind := i  - i%3 + m
						if ind == i {
							continue
						}
						removed = g.Rows[i  - i%3 + m][k*3+l].RemoveCandidateWithFeedback(val) || removed
					}
				}

				if removed {
					box := k + (i / 3)*3
					fmt.Printf("box/row reduction of %d in box %d, row %d\n", val, box+1, i+1)
					performedAction = true
				}
			}
		}
	}

	// Columns
	for i := 0; i < 9; i++ {
		for val := 1; val <= 9; val++ {
			posOnCol:= g.DigitsPositions.Cols[i][val-1]
			if posOnCol == 0 { // digit already in row
				continue
			}

			for k, a := range alignments {
				if posOnCol&a != posOnCol { // values are not on the current alignment mask
					continue
				}

				
				removed := false
				for l := 0; l < 3; l++ {	
					for m := 0; m < 3; m++ {
						ind := i  - i%3 + m
						if ind == i {
							continue
						}
						removed = g.Cols[i  - i%3 + m][k*3+l].RemoveCandidateWithFeedback(val) || removed
					}
				}

				if removed {
					box := i/3 + k*3
					fmt.Printf("box/col reduction of %d in box %d, col %d\n", val, box+1, i+1)
					performedAction = true
				}
			}
		}
	}

	return performedAction
}
