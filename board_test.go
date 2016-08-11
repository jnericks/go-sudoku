package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddRow(t *testing.T) {
	Convey("When adding a row to a board", t, func() {

		board := NewBoard()
		board.AddRow(0, "1,2,-,4,5,6,-,8,9")
		Convey("it should parse the row correctly into cells", func() {
			So(board.Cells[Coordinate{0, 0}].Solution, ShouldEqual, 1)
			So(board.Cells[Coordinate{0, 1}].Solution, ShouldEqual, 2)
			So(board.Cells[Coordinate{0, 2}].Solution, ShouldEqual, 0)
			So(board.Cells[Coordinate{0, 3}].Solution, ShouldEqual, 4)
			So(board.Cells[Coordinate{0, 4}].Solution, ShouldEqual, 5)
			So(board.Cells[Coordinate{0, 5}].Solution, ShouldEqual, 6)
			So(board.Cells[Coordinate{0, 6}].Solution, ShouldEqual, 0)
			So(board.Cells[Coordinate{0, 7}].Solution, ShouldEqual, 8)
			So(board.Cells[Coordinate{0, 8}].Solution, ShouldEqual, 9)
		})

		Convey("it should eliminate possible values", func() {
			cell := board.Cells[Coordinate{1, 0}]
			So(cell.Possible, ShouldNotContain, 1)
			So(cell.Possible, ShouldNotContain, 2)
			So(cell.Possible, ShouldContain, 3)
			So(cell.Possible, ShouldContain, 4)
			So(cell.Possible, ShouldContain, 5)
			So(cell.Possible, ShouldContain, 6)
			So(cell.Possible, ShouldContain, 7)
			So(cell.Possible, ShouldContain, 8)
			So(cell.Possible, ShouldContain, 9)
		})
	})
}
