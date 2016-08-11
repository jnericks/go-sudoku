package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSample(t *testing.T) {
	Convey("When solving a valid board", t, func() {
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

		board, _ := br.ReadBoard()
		var solved bool
		board, solved = solve(board)

		Convey("it should be reported as solved", func() {
			So(solved, ShouldBeTrue)
		})
	})
}
