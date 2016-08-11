package main

import (
	"os"
	"path"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// GetDataFile will return a full path to a file in the /data directory
func GetDataFile(file string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Join(dir, "/data", file), nil
}

func TestGetDataFile(t *testing.T) {
	Convey("When getting a full path for a sudoku csv test file", t, func() {
		fullpath, _ := GetDataFile("sudoku1.csv")

		Convey("then it should be a file that exists", func() {
			fi, err := os.Stat(fullpath)
			So(err, ShouldBeNil)
			So(fi.IsDir(), ShouldBeFalse)
		})

		Convey("then the file should be in the /data directory", func() {
			So(fullpath, ShouldEndWith, "/data/sudoku1.csv")
		})
	})
}

func TestCsvBoardReader(t *testing.T) {
	Convey("When reading a board from sudoku1.csv", t, func() {
		var br BoardReader
		file, _ := GetDataFile("sudoku1.csv")
		br = CsvBoardReader{file}
		board, err := br.ReadBoard()

		Convey("it should not have an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("it should have a board with 81 cells (9x9)", func() {
			So(len(board.Cells), ShouldEqual, 81)
		})
	})
}
