package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrans(t *testing.T) {
	Convey("uint to str", t, func() {

		Convey("0", func() {
			s := trans(0)
			So(s, ShouldEqual, "000000")
		})

		Convey("1", func() {
			s := trans(1)
			So(s, ShouldEqual, "000001")
		})

		Convey("10", func() {
			s := trans(10)
			So(s, ShouldEqual, "00000a")
		})

		Convey("100", func() {
			s := trans(100)
			So(s, ShouldEqual, "00001C")
		})

		Convey("1000", func() {
			s := trans(1000)
			So(s, ShouldEqual, "0000g8")
		})
		Convey("10000", func() {
			s := trans(10000)
			So(s, ShouldEqual, "0002Bi")
		})

		Convey("100000", func() {
			s := trans(100000)
			So(s, ShouldEqual, "000q0U")
		})

		Convey("1000000", func() {
			s := trans(1000000)
			So(s, ShouldEqual, "004c92")
		})

		Convey("10000000", func() {
			s := trans(10000000)
			So(s, ShouldEqual, "00FXsk")
		})

		Convey("100000000", func() {
			s := trans(100000000)
			So(s, ShouldEqual, "06LAze")
		})

		Convey("1000000000", func() {
			s := trans(1000000000)
			So(s, ShouldEqual, "15FTGg")
		})

	})
}
