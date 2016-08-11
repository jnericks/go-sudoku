package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNotifyUpdate(t *testing.T) {
	Convey("When notifing a cell that a cell in the same row is being updated", t, func() {

		cell := UnsolvedCell(Coordinate{0, 0})
		var adjusted bool
		cell, adjusted = cell.NotifyUpdate(SolvedCell(Coordinate{1, 0}, 1))
		Convey("it should reduce the possible values", func() {
			So(adjusted, ShouldBeTrue)
			So(cell.Possible, ShouldNotContain, 1)
			So(cell.Possible, ShouldContain, 2)
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
