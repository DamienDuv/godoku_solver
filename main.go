package main

import (
	"fmt"

	"github.com/DamienDuv/godoku_solver/models"
	"github.com/DamienDuv/godoku_solver/solver"
)


func main() {
	
	g, err := models.NewGrid(solver.Expert[2])

	if err != nil {
		fmt.Printf("%v\n",err)
		return
	}

	solver.Solve(g)

}