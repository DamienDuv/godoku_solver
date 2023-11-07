package models

import (
	"fmt"
)

type DigitsPositionsMasks struct {
	Rows  [][]uint16
	Cols  [][]uint16
	Boxes [][]uint16
}

type Grid struct {
	Grid []Cell
	Rows [][]*Cell
	Cols [][]*Cell
	Boxes [][]*Cell

	DigitsPositions DigitsPositionsMasks
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
		ok := isSetValid(g.Rows[i])
		if !ok {
			fmt.Printf("Row %d\n",i+1)
			return false
		}

		ok = isSetValid(g.Cols[i])
		if !ok {
			fmt.Printf("Col %d\n",i+1)
			return false
		}

		ok = isSetValid(g.Boxes[i])
		if !ok {
			fmt.Printf("Box %d\n",i+1)
			return false
		}
	}

	for i := 0; i < 81;i++ {
		if g.Grid[i].HasNoCandidate() {
			r,c := IndexToCoord(i)
			fmt.Printf("r%dc%d has no candidates\n",r+1,c+1)
			return false 
		}
	}

	return true
}

func (g Grid) IsFilled() bool {
	for _,v := range g.Grid {
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
	g.Rows[row][col].setValue(n)
	
	box := CoordToBox(row,col)
	for i := 0; i < 9; i++ {
		g.Rows[row][i].RemoveCandidate(n)
		g.Cols[col][i].RemoveCandidate(n)
		g.Boxes[box][i].RemoveCandidate(n)
	}
}

// UpdateDigitsPositionMask creates a map with 3 keys (rows,cols,boxes)
// each element contains a 2D array of uint16. The first dimension correspond to the set number (row 1, col3...)
// the second dimension correspond to a sudoku number (base 0) for witch you want the mask in the row
// ex: row[0][4] = 0b100100100 means that in row 1, the 5 could be in positions 3,6 or 9
func (g * Grid) UpdateDigitsPositionMask() {
	for c := 1; c <= 9; c++ {
		var mask uint16 = 1 << (c - 1)
		for i := 0; i < 9; i++ {
			g.DigitsPositions.Rows[i][c-1] = 0
			g.DigitsPositions.Cols[i][c-1] = 0
			g.DigitsPositions.Boxes[i][c-1] = 0
			for j := 0; j < 9; j++ {

				cBitRow := mask & g.Rows[i][j].GetCandidates()
				cBitCol := mask & g.Cols[i][j].GetCandidates()
				cBitBox := mask & g.Boxes[i][j].GetCandidates()

				if j-(c-1) < 0 {
					g.DigitsPositions.Rows[i][c-1] = g.DigitsPositions.Rows[i][c-1] | (cBitRow >> (c - 1 - j))
					g.DigitsPositions.Cols[i][c-1] = g.DigitsPositions.Cols[i][c-1] | (cBitCol >> (c - 1 - j))
					g.DigitsPositions.Boxes[i][c-1] = g.DigitsPositions.Boxes[i][c-1] | (cBitBox >> (c - 1 - j))
				} else {
					g.DigitsPositions.Rows[i][c-1] = g.DigitsPositions.Rows[i][c-1] | (cBitRow << (j - c + 1))
					g.DigitsPositions.Cols[i][c-1] = g.DigitsPositions.Cols[i][c-1] | (cBitCol << (j - c + 1))
					g.DigitsPositions.Boxes[i][c-1] = g.DigitsPositions.Boxes[i][c-1] | (cBitBox << (j - c + 1))
				}

			}
		}
	}
}

func newEmptyGrid() *Grid {
	g := &Grid{}
	g.Grid = make([]Cell, 81)
	for i := range g.Grid {
		g.Grid[i] = NewCell()
	}
	g.Rows = make([][]*Cell,9)
	g.Cols = make([][]*Cell,9)
	g.Boxes = make([][]*Cell,9)
	g.DigitsPositions.Rows = make([][]uint16, 9)
	g.DigitsPositions.Cols = make([][]uint16, 9)
	g.DigitsPositions.Boxes = make([][]uint16, 9)


	for i := 0; i < 9; i++ {
		g.Rows[i] = make([]*Cell, 9)
		g.Cols[i] = make([]*Cell, 9)
		g.Boxes[i] = make([]*Cell, 9)
		g.DigitsPositions.Rows[i] = make([]uint16, 9)
		g.DigitsPositions.Cols[i] = make([]uint16, 9)
		g.DigitsPositions.Boxes[i] = make([]uint16, 9)

		for j := 0; j < 9; j++ {
			g.Rows[i][j] = &(g.Grid[i*9 + j])
			g.Cols[i][j] = &(g.Grid[i + j*9])
			g.Boxes[i][j] = &(g.Grid[j%3 + (j/3)*9 + 3*i + (i/3)*18])
		}
	}
	return g
}

func (g Grid) Print() {
	fmt.Print("|-----------------------------|\n")
	for i := 0; i < 9; i++ {
		fmt.Print("|")
		for j := 0 ; j < 9; j++ {
			v := g.Grid[9*i+j].val

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

func (g Grid) String() string{
	bytes := make([]byte,81)

	for i,v := range g.Grid {
		if v.val == 0 {
			bytes[i] = '.'
		} else {
			bytes[i] = byte(v.val) +'0'
		}
	}

	return string(bytes)
}