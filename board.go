package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Board represents a sudoku game board
type Board struct {
	Cells map[Coordinate]Cell
}

// NewBoard will construct a 9x9 Sudoku Board with unsolved cells
func NewBoard() Board {
	cells := make(map[Coordinate]Cell)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			c := Coordinate{row, col}
			cells[c] = UnsolvedCell(c)
		}
	}
	return Board{cells}
}

// IsSolved will report if the sudoku board was solved
func (b Board) IsSolved() bool {
	for _, c := range b.Cells {
		if !c.IsSolved() {
			return false
		}
	}
	return true
}

// IsSolvable will report if the sudoku board has the potential to still be solved
func (b Board) IsSolvable() bool {
	for _, c := range b.Cells {
		if !c.IsSolved() && !c.HasPossibleSolutions() {
			return false
		}
	}
	return true
}

// AddRow will validate and then store the row of the board
func (b Board) AddRow(row int, input string) (Board, error) {
	if b.Cells == nil {
		return b, RaiseError("Board is not initialized, Cells is nil")
	}

	if row < 0 || row > 8 {
		return b, RaiseError("Row should be a value in-between 0 and 8 (inclusive)")
	}

	values := strings.Split(input, ",")
	//fmt.Printf("Row: %d, Input: '%s', Values: %s\n", row, input, values)
	if len(values) != 9 {
		return b, RaiseErrorf("Expecting 9 values comma-separated per sudoku row, row: %d, input: '%s'", row, input)
	}

	for col, v := range values {
		coordinate := Coordinate{row, col}

		switch v {
		case "-", "0":
			// do nothing
			//fmt.Printf("{%d, %d}: do nothing\n", c.Row, c.Column)
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			solution, _ := strconv.Atoi(v)
			var err error
			b, err = b.SetSolution(coordinate, solution)
			if err != nil {
				return b, err
			}
		default:
			return b, RaiseErrorf("Invalid cell, row: %d, col: %d value: %s", coordinate.Row, coordinate.Column, v)
		}
	}

	return b, nil
}

// SetSolution sets the solution of the cell at the coordinate
func (b Board) SetSolution(coordinate Coordinate, solution int) (Board, error) {
	//log.Printf("{%d, %d}: b.SetSolution({%d, %d}, %d)", coordinate.Row, coordinate.Column, coordinate.Row, coordinate.Column, solution)

	if b.Cells[coordinate].IsSolved() {
		return b, nil
	}

	found := false
	for _, value := range b.Cells[coordinate].Possible {
		if found = solution == value; found {
			break
		}
	}
	if !found {
		return b, RaiseErrorf("%d is not a possible solution for cell {%d, %d}, possible: %v", solution, coordinate.Row, coordinate.Column, b.Cells[coordinate].Possible)
	}

	solved := SolvedCell(coordinate, solution)
	b.Cells[coordinate] = solved

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			c := b.Cells[Coordinate{row, col}]
			var adjusted bool
			c, adjusted = c.NotifyUpdate(solved)
			b.Cells[Coordinate{row, col}] = c
			if adjusted && len(c.Possible) == 1 {
				var err error
				if b, err = b.SetSolution(c.Coordinate, c.Possible[0]); err != nil {
					return b, err
				}
			}
		}
	}

	return b, nil
}

func (b Board) String() string {
	var buffer bytes.Buffer

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			buffer.WriteString(fmt.Sprintf("%d", b.Cells[Coordinate{row, col}].Solution))
			if col != 8 {
				buffer.WriteString(",")
			}
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}
