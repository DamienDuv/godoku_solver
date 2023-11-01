package main

type Cell struct {
	val int
	candidates uint16
}

// NewCell creates an empty Cell with all possible candidates
func NewCell() Cell {
	return Cell{candidates: 0b111111111}
}

// RemoveCandidate removes the candiates {n} from the possible candidates
func (c *Cell) RemoveCandidate(n int) {
	c.candidates = c.candidates &^ (1<< (n-1))
}

// HasCandidate checks if the Cell has the candidate {n}
func (c *Cell) HasCandidate(n int) bool {
	return c.candidates & (1<< (n-1)) != 0
}

// HasNoCandidates returns true if there is no possible candidate for this Cell (i.e invalid grid)
func (c *Cell) HasNoCandidate() bool {
	return c.candidates == 0
}

// IsSet return true if the Cell has a value and false otherwise
func (c Cell) IsSet() bool {
	return c.val != 0
}

// HasOneCandidate check if a Cell has only one candidate possible
// and returns it if it exists
func (c *Cell) HasOneCandidate() (bool, int) {
	for i := 0; i < 9; i++ {
		var mask uint16 = (1 << i)
		if (c.candidates & mask != 0) && (c.candidates ^ mask == 0) {
			return true, i + 1
		}
	}
	return false, 0
}

// IndexToCoord turns an index in [0, 80] to its sudoku grid
// row , col coordinates in [0, 8]
func IndexToCoord(i int) (int,int) {
	row := i / 9
	col := i % 9

	return row, col
}

// CoordToBox turns sudoku grid coordinates into the coresponding block number
func CoordToBox(r int, c int) int {
	return (r/3)*3 + c / 3
}
