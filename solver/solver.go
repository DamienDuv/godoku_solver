package solver

import (
	"fmt"

	"github.com/DamienDuv/godoku/models"
)

const limit = 50

func Solve(g *models.Grid) {
	i := 0
	performedAction := true

	for performedAction && i < limit{
		fmt.Println(i+1, " ===============================")
		performedAction = false
		performedAction = performedAction || nakedSingles(g)

		g.UpdateDigitsPositionMask()
		performedAction = performedAction || hiddenSingle(g)

		if g.IsFilled() && g.IsValid() {
			g.Print()
			fmt.Printf("Solved in %d iterations\n", i+1)
			return
		}

		performedAction = performedAction || alignmentInBox(g)

		for j := 2; j <= 4; j++ {
			performedAction = performedAction || nakedSubSet(g, j)
			performedAction = performedAction || hiddenSubSet(g, j)
		}

		i++
	}


	fmt.Printf("%b\n",g.DigitsPositions.Boxes[0])
	

	g.Print()
	fmt.Printf("Couldn't solve the grid any further with current settings\n")
}