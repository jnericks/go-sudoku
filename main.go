package main

import (
	"fmt"
	"os"
)

func main() {

	var br BoardReader
	args := os.Args
	if len(args) > 1 {
		br = CsvBoardReader{args[1]}
	} else {
		br = StdinBoardReader{}
	}

	//var br BoardReader = CsvBoardReader{"hardest.csv"}

	/*
		var br BoardReader = InMemoryBoardReader{[]string{
			"-,-,3,-,2,-,6,-,-",
			"9,-,-,3,-,5,-,-,1",
			"-,-,1,8,-,6,4,-,-",
			"-,-,8,1,-,2,9,-,-",
			"7,-,-,-,-,-,-,-,8",
			"-,-,6,7,-,8,2,-,-",
			"-,-,2,6,-,9,5,-,-",
			"8,-,-,2,-,3,-,-,9",
			"-,-,5,-,1,-,3,-,-"}}
	*/

	board, err := br.ReadBoard()

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	var solved bool
	board, solved = solve(board)

	if !solved {
		fmt.Println("No solution found")
		os.Exit(0)
	}

	fmt.Println("Solution:")
	fmt.Println(board.String())
}

func solve(b Board) (Board, bool) {
	if b.IsSolved() {
		return b, true
	}

	for b.IsSolvable() {
		for _, c := range b.Cells {
			if !c.IsSolved() {
				var err error
				b, err = b.SetSolution(c.Coordinate, c.Possible[0])
				if err != nil {
					return b, false
				}
				if b.IsSolved() {
					return b, true
				}
				return solve(b)
			}
		}
	}

	return b, false
}
