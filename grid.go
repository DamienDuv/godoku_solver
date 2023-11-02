package main

import (
	"fmt"
)

type Grid struct {
	grid []Cell
	rows [][]*Cell
	cols [][]*Cell
	boxes [][]*Cell
}

func NewGrid(s string) (*Grid, error) {
	if len(s) != 81 {
		return nil, fmt.Errorf("error parsing input grid string, wrong size : %s", s)
	}

	g := newEmptyGrid()

	// Fill the grid
	for i,v := range s {
		if v >= '1' && v <= '9' {
			g.InsertNumberByIndex(int(v - '1' + 1),i)
			continue
		} else if v != '.' {
			return nil, fmt.Errorf("error parsing input grid string, unsupported character : %s", s)
		}
	}

	return g, nil
}

func (g Grid) IsValid() bool {
	for i := 0; i < 9; i++ {
		ok := isSetValid(g.rows[i])
		if !ok {
			fmt.Printf("Row %d\n",i+1)
			return false
		}

		ok = isSetValid(g.cols[i])
		if !ok {
			fmt.Printf("Col %d\n",i+1)
			return false
		}

		ok = isSetValid(g.boxes[i])
		if !ok {
			fmt.Printf("Box %d\n",i+1)
			return false
		}
	}

	for i := 0; i < 81;i++ {
		if g.grid[i].HasNoCandidate() {
			r,c := IndexToCoord(i)
			fmt.Printf("r%dc%d has no candidates\n",r+1,c+1)
			return false 
		}
	}

	return true
}

func (g Grid) IsFilled() bool {
	for _,v := range g.grid {
		if !v.IsSet() {
			return false
		}
	}
	return true
}

func isSetValid(set []*Cell) bool {
	present := make([]bool,9)
	for i:= 0; i < 9; i++ {
		n := set[i].val
		if n != 0 {
			if present[n-1] {
				return false
			}

			present[n-1] = true
		}
	}

	return true
}

func (g *Grid) InsertNumberByIndex(n int, index int) {
	row, col := IndexToCoord(index)
	g.InsertNumberByCoord(n,row,col)
}

func (g *Grid) InsertNumberByCoord(n int, row int , col int) {
	g.rows[row][col].SetValue(n)
	
	box := CoordToBox(row,col)
	for i := 0; i < 9; i++ {
		g.rows[row][i].RemoveCandidate(n)
		g.cols[col][i].RemoveCandidate(n)
		g.boxes[box][i].RemoveCandidate(n)
	}
}

func newEmptyGrid() *Grid {
	g := &Grid{}
	g.grid = make([]Cell, 81)
	for i := range g.grid {
		g.grid[i] = NewCell()
	}

	g.rows = make([][]*Cell,9)
	g.cols = make([][]*Cell,9)
	g.boxes = make([][]*Cell,9)

	for i := 0; i < 9; i++ {
		g.rows[i] = make([]*Cell, 9)
		g.cols[i] = make([]*Cell, 9)
		g.boxes[i] = make([]*Cell, 9)

		for j := 0; j < 9; j++ {
			g.rows[i][j] = &(g.grid[i*9 + j])
			g.cols[i][j] = &(g.grid[i + j*9])
			g.boxes[i][j] = &(g.grid[j%3 + (j/3)*9 + 3*i + (i/3)*18])
		}
	}
	return g
}

func (g Grid) Print() {
	fmt.Print("|-----------------------------|\n")
	for i := 0; i < 9; i++ {
		fmt.Print("|")
		for j := 0 ; j < 9; j++ {
			v := g.grid[9*i+j].val

			if v != 0 {
				fmt.Print(" ",v," ")
			} else {
				fmt.Print("   ")
			}

			if (j+1) %3 == 0 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
		if (i + 1)%3 == 0{
			fmt.Print("|-----------------------------|\n")
		}
	}
}