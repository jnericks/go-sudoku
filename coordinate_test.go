package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSectory(t *testing.T) {
	Convey("When getting the sector", t, func() {
		Convey("it be able to resolve each sector", func() {
			So(Coordinate{0, 0}.GetSector(), ShouldEqual, 0)
			So(Coordinate{3, 1}.GetSector(), ShouldEqual, 1)
			So(Coordinate{6, 2}.GetSector(), ShouldEqual, 2)
			So(Coordinate{1, 3}.GetSector(), ShouldEqual, 3)
			So(Coordinate{4, 4}.GetSector(), ShouldEqual, 4)
			So(Coordinate{7, 5}.GetSector(), ShouldEqual, 5)
			So(Coordinate{2, 6}.GetSector(), ShouldEqual, 6)
			So(Coordinate{5, 7}.GetSector(), ShouldEqual, 7)
			So(Coordinate{8, 8}.GetSector(), ShouldEqual, 8)
		})
	})
}
