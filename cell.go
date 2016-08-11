package main

// Cell represents a single cell in sudoku
type Cell struct {
	Coordinate Coordinate
	Solution   int
	Possible   []int
}

// UnsolvedCell returns an unsolved cell
func UnsolvedCell(coordinate Coordinate) Cell {
	return Cell{Coordinate: coordinate, Solution: 0, Possible: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
}

// SolvedCell returns a solved cell
func SolvedCell(coordinate Coordinate, solution int) Cell {
	return Cell{Coordinate: coordinate, Solution: solution, Possible: []int{}}
}

// IsSolved will report if a cell has been solved
func (c Cell) IsSolved() bool {
	return c.Solution > 0
}

// HasPossibleSolutions will report if a cell has possible solutions
func (c Cell) HasPossibleSolutions() bool {
	return len(c.Possible) > 0
}

// NotifyUpdate ...
func (c Cell) NotifyUpdate(other Cell) (Cell, bool) {
	if c.Coordinate.Row == other.Coordinate.Row || // same row
		c.Coordinate.Column == other.Coordinate.Column || // same column
		c.Coordinate.GetSector() == other.Coordinate.GetSector() { // same sector
		var i int
		var found bool
		var v int
		for i, v = range c.Possible {
			if v == other.Solution {
				found = true
				break
			}
		}
		if found {
			// reduce possible solutions
			c.Possible = append(c.Possible[:i], c.Possible[i+1:]...)
			return c, true
		}
	}
	return c, false
}
