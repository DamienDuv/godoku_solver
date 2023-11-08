package godoku_solver

import (
	"github.com/DamienDuv/godoku_solver/models"
	"github.com/DamienDuv/godoku_solver/solver"
)

func SolveSudoku(sudokuString string) (solution string, err error) {
	g, err := models.NewGrid(sudokuString)
	if err != nil {
		return "", err
	}

	solver.Solve(g)

	return g.String(), nil
}