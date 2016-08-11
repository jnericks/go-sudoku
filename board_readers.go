package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// BoardReader will read in a sudoku board
type BoardReader interface {
	ReadBoard() (Board, error)
}

// StdinBoardReader will read in a sudoku board from Stdin
type StdinBoardReader struct{}

// CsvBoardReader will read in a sudoku board from the given file
type CsvBoardReader struct {
	File string
}

// InMemoryBoardReader will read in a sudoku board from the given lines
type InMemoryBoardReader struct {
	Lines []string
}

// ReadBoard will read in a sudoku board from Stdin
func (br StdinBoardReader) ReadBoard() (Board, error) {
	board := NewBoard()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Sudoku:")

	row := 0
	for scanner.Scan() {
		for _, input := range strings.Split(scanner.Text(), "\n") {
			var err error
			if board, err = board.AddRow(row, input); err != nil {
				return board, err
			}
			row++
		}
		if row > 8 {
			break
		}
	}

	return board, nil
}

// ReadBoard will read in a sudoku board from the file specified in the CsvBoardReader struct
func (br CsvBoardReader) ReadBoard() (Board, error) {
	board := NewBoard()

	var bytes []byte
	bytes, err := ioutil.ReadFile(br.File)
	if err != nil {
		return board, err
	}

	for row, input := range strings.Split(string(bytes), "\n") {
		if row > 8 {
			break
		}
		fmt.Println(input)
		if board, err = board.AddRow(row, input); err != nil {
			return board, err
		}
	}

	return board, nil
}

// ReadBoard will read in a sudoku board from the lines specified in the InMemoryBoardReader struct
func (br InMemoryBoardReader) ReadBoard() (Board, error) {
	board := NewBoard()
	for row, input := range br.Lines {
		var err error
		if board, err = board.AddRow(row, input); err != nil {
			return board, err
		}
	}
	return board, nil
}
