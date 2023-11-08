package solver

import (
	"fmt"

	"github.com/DamienDuv/godoku_solver/models"
)

// Find Cells where there is only one candidate left and fills it
func nakedSingles(g *models.Grid) bool {
	performedAction := false
	for i := range g.Grid {
		n, ok := g.Grid[i].HasOneCandidate()
		if ok {
			performedAction = true
			g.InsertNumberByIndex(n, i)
			r, c := models.IndexToCoord(i)
			fmt.Printf("naked single r%dc%d = %d\n", r+1, c+1, n)
		}
	}
	return performedAction
}