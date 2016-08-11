package main

// Coordinate represents a coordinate on the sudoku board
type Coordinate struct {
	Row, Column int
}

// GetSector will get the sector of a Coordinate
func (c Coordinate) GetSector() int {
	if c.Column <= 2 {
		if c.Row <= 2 {
			return 0
		}
		if c.Row <= 5 {
			return 1
		}
		if c.Row <= 8 {
			return 2
		}
		return -1
	}
	if c.Column <= 5 {
		if c.Row <= 2 {
			return 3
		}
		if c.Row <= 5 {
			return 4
		}
		if c.Row <= 8 {
			return 5
		}
		return -1
	}
	if c.Column <= 8 {
		if c.Row <= 2 {
			return 6
		}
		if c.Row <= 5 {
			return 7
		}
		if c.Row <= 8 {
			return 8
		}
		return -1
	}
	return -1
}
