package main

type DigitsPositionsMasks struct {
	rows [][]uint16
	cols [][]uint16
	boxes [][]uint16
}

// BuildNumberCandidatesPosMasks creates a map with 3 keys (rows,cols,boxes)
// each element contains a 2D array of uint16. The first dimension correspond to a sudoku number
// the second dimension
func NewDigitsPositionsMasks(g Grid) DigitsPositionsMasks {
	p := DigitsPositionsMasks{}
	p.rows = make([][]uint16, 9)
	p.cols= make([][]uint16, 9)
	p.boxes = make([][]uint16, 9)

	for c := 1; c <= 9; c++ {
		var mask uint16 = 1 << (c - 1)
		p.rows[c-1] = make([]uint16, 9)
		p.cols[c-1] = make([]uint16, 9)
		p.boxes[c-1] = make([]uint16, 9)

		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				
				cBitRow := mask & g.rows[i][j].candidates
				cBitCol := mask & g.cols[i][j].candidates
				cBitBox := mask & g.boxes[i][j].candidates

				if j-(c-1) < 0 {
					p.rows[c-1][i] = p.rows[c-1][i] | (cBitRow >> (c - 1 - j))
					p.cols[c-1][i] = p.cols[c-1][i] | (cBitCol >> (c - 1 - j))
					p.boxes[c-1][i] = p.boxes[c-1][i] | (cBitBox >> (c - 1 - j))
				} else {
					p.rows[c-1][i] = p.rows[c-1][i] | (cBitRow << (j - c + 1))
					p.cols[c-1][i] = p.cols[c-1][i] | (cBitCol << (j - c + 1))
					p.boxes[c-1][i] = p.boxes[c-1][i] | (cBitBox << (j - c + 1))
				}

			}
		}
	}

	return p
}
