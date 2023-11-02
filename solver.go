package main

import (
	"fmt"
	"math/bits"
)

func Solve(g *Grid) {
	i := 0
	performedAction := true

	for performedAction {
		fmt.Println(i + 1, " ===============================")
		performedAction = false
		performedAction = performedAction || nakedSingles(g)

		pos := NewDigitsPositionsMasks(*g)
		performedAction = performedAction || isolatedSingle(g, pos)

		if g.IsFilled() && g.IsValid() {
			g.Print()
			fmt.Printf("Solved in %d iterations\n", i+1)
			return
		}

		performedAction = performedAction || alignmentInBox(g, pos)
		performedAction = performedAction || hiddenPair(g, pos)

		i++
	}

	g.Print()
	fmt.Printf("Couldn't solve the grid any further with current settings\n")
}

// Find Cells where there is only one candidate left and fills it
func nakedSingles(g *Grid) bool {
	performedAction := false
	for i := range g.grid {
		n, ok := g.grid[i].HasOneCandidate()
		if ok {
			performedAction = true
			g.InsertNumberByIndex(n, i)
			r, c := IndexToCoord(i)
			fmt.Printf("nakedSingles : r%dc%d = %d\n", r+1, c+1, n)
		}
	}
	return performedAction
}

// Find Rows / Columns / Boxes where a candidate is present only once and fills it
func isolatedSingle(g *Grid, pos DigitsPositionsMasks) bool {
	performedAction := false
	for val := 1; val <= 9; val++ {
		for i := 0; i < 9; i++ {
			n, ok := IsolatedBitIndex(pos.rows[val-1][i])
			if ok {
				fmt.Printf("[row] isolated single %d at r%dc%d\n", val, i+1, n)
				g.InsertNumberByCoord(val, i, n-1)
				performedAction = true
			}

			n, ok = IsolatedBitIndex(pos.cols[val-1][i])
			if ok {
				fmt.Printf("[col] isolated single %d at r%dc%d\n", val, n, i+1)
				g.InsertNumberByCoord(val, n-1, i)
				performedAction = true
			}

			n, ok = IsolatedBitIndex(pos.boxes[val-1][i])
			if ok {
				fmt.Printf("[box] isolated single %d in box %d, pos %d\n", val, i+1, n)
				g.InsertNumberByIndex(val, BoxCoordToGridIndex(i+1, n))
				performedAction = true
			}
		}
	}
	return performedAction
}

// Finds an alignment of the same candidate withing a box and clears the corresponding
// row or column of this candidate (outside of the box containing the alignment)
func alignmentInBox(g *Grid, pos DigitsPositionsMasks) bool {
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
			posInBox := pos.boxes[val-1][i]
			if posInBox == 0 { // digit already in box
				continue
			}

			for k, a := range alignments {
				if posInBox & a != posInBox { // values are not on the current alignment mask
					continue
				}

				if k < 3 { // is on a row alignment
					row := (i/3)*3 + k
					removed := false
					for l := 0; l < 9; l++ {
						if l < (i%3)*3 || l >= (i%3)*3+3 {
							res := g.rows[row][l].RemoveCandidateWithFeedback(val)
							if res {
								fmt.Printf("cleared marks %d on r%dc%d\n",val,row+1,l+1)
							}
							removed = removed || res
						}
					}

					if removed  {
						fmt.Printf("Alignment of %d in box %d, row %d\n", val, i+1, row+1)
						performedAction = true
					}
				} else { // is on a column alignment
					col := (i%3)*3 + k - 3
					removed := false
					for l := 0; l < 9; l++ {
						if l < (i/3)*3 || l >= (i/3)*3+3 {
							res := g.cols[col][l].RemoveCandidateWithFeedback(val)
							if res {
								fmt.Printf("cleared marks %d on r%dc%d\n",val,l+1,col+1)
							}
							removed = removed || res
						}
					}

					if removed  {
						fmt.Printf("Alignment of %d in box %d, col %d\n", val, i+1, col+1)
						performedAction = true
					}

				}
				break

			}
		}
	}

	return performedAction
}

func hiddenPair(g *Grid, pos DigitsPositionsMasks) bool {
	performedAction := false
	for i := 0; i < 9; i++ {
		// look for two digits v and d having the same mask on the row
		// and have only two occurences. if we found one, it is a pair
		for v := 1; v <= 9; v++ {
			for d := v+1; d <= 9; d++ {
				if pos.rows[v-1][i] != pos.rows[d-1][i] {
					continue
				}

				if bits.OnesCount16(pos.rows[v-1][i]) != 2 {
					continue
				}
				if bits.OnesCount16(pos.rows[d-1][i]) != 2 {
					continue
				}

				// we have a pair (no pun intended)
				indexes := BitsIndexes(pos.rows[v-1][i])

				removed := false
				for _, ind := range indexes {
					res := g.rows[i][ind].RemoveCandidateExcept(v,d)
					removed = removed || res
				}

				if removed {
					performedAction = true
					fmt.Printf("found a %d/%d pair on row %d at col %d/%d\n",v,d,i+1,indexes[0]+1,indexes[1]+1)
				}

				

			}
		}
	}
	return performedAction
}