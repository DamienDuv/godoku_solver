package main

type Cell struct {
	val        int
	candidates uint16
}

// NewCell creates an empty Cell with all possible candidates
func NewCell() Cell {
	return Cell{candidates: 0b111111111}
}

// SetValue sets a cell value and remove all candidates
func (c *Cell) SetValue(v int) {
	c.val = v
	c.candidates = 0
}

// RemoveCandidate removes the candiates {n} from the possible candidates
func (c *Cell) RemoveCandidate(n int) {
	c.candidates = c.candidates &^ (1 << (n - 1))
}

// RemoveCandidateWithFeedback removes the candiates {n} from the possible candidates
// returs true if a value was removed, false if the value was already removed
func (c *Cell) RemoveCandidateWithFeedback(n int) bool {
	if !c.HasCandidate(n) {
		return false
	}

	c.candidates = c.candidates &^ (1 << (n - 1))
	return true
}

// RemoveCandidateExcept removes all candidates except the ones listed in
// the arguments
func (c *Cell) RemoveCandidateExcept(digits ...int) bool {
	mask := uint16(0)
	for _, n := range digits {
		mask = mask | (1 << (n - 1))
	}
	mask = ^mask

	newCandidates := c.candidates &^ mask

	if newCandidates == c.candidates {
		return false
	}

	c.candidates = newCandidates
	return true
}

// HasCandidate checks if the Cell has the candidate {n}
func (c *Cell) HasCandidate(n int) bool {
	return c.candidates&(1<<(n-1)) != 0
}

// HasNoCandidates returns true if there is no possible candidate for this Cell (i.e invalid grid)
func (c *Cell) HasNoCandidate() bool {
	return c.candidates == 0 && c.val == 0
}

// IsSet return true if the Cell has a value and false otherwise
func (c Cell) IsSet() bool {
	return c.val != 0
}

// HasOneCandidate check if a Cell has only one candidate possible
// and returns it if it exists
func (c Cell) HasOneCandidate() (int, bool) {
	return IsolatedBitIndex(c.candidates)
}

// IndexToCoord turns an index in [0, 80] to its sudoku grid
// row , col coordinates in [0, 8]
func IndexToCoord(i int) (int, int) {
	row := i / 9
	col := i % 9

	return row, col
}

// CoordToBox turns sudoku grid coordinates into the coresponding block number
func CoordToBox(r int, c int) int {
	return (r/3)*3 + c/3
}

func BoxCoordToGridIndex(box int, pos int) int {
	start := ((box-1)%3)*3 + ((box-1)/3)*27
	return start + (pos-1)%3 + ((pos-1)/3)*9
}
