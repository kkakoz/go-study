package goconvey_study

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual should return true when a != nil && b != nil",
		t, func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeTrue)
		},
	)

	Convey("TestStringSliceEqual should return false a != b",
		t, func() {
			a := []string{"hello", "goc"}
			b := []string{"hello", "go"}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		},
	)

	Convey("equal", t, func() {
		So(2.0, ShouldEqual, 2)
	})
}
