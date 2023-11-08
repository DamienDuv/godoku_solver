package solver

import (
	"fmt"

	"github.com/DamienDuv/godoku_solver/models"
)

// Find Rows / Columns / Boxes where a candidate is present only once and fills it
func hiddenSingle(g *models.Grid) bool {
	performedAction := false
	for val := 1; val <= 9; val++ {
		for i := 0; i < 9; i++ {
			n, ok := models.IsolatedBitIndex(g.DigitsPositions.Rows[i][val-1])
			if ok {
				fmt.Printf("hidden single %d in row %d at col %d\n", val, i+1, n)
				g.InsertNumberByCoord(val, i, n-1)
				performedAction = true
			}

			n, ok = models.IsolatedBitIndex(g.DigitsPositions.Cols[i][val-1])
			if ok {
				fmt.Printf("hidden single %d in col %d at row %d\n", val, i+1,n)
				g.InsertNumberByCoord(val, n-1, i)
				performedAction = true
			}

			n, ok = models.IsolatedBitIndex(g.DigitsPositions.Boxes[i][val-1])
			if ok {
				fmt.Printf("hidden single %d in box %d, pos %d\n", val, i+1, n)
				g.InsertNumberByIndex(val, models.BoxCoordToGridIndex(i+1, n))
				performedAction = true
			}
		}
	}
	return performedAction
}