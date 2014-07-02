package message

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMessageTypeLookup(t *testing.T) {
	Convey("Various strings should map to a MessageType when using Lookup()", t, func() {
		Convey("'error' should map to ERROR_MESSAGE", func() {
			So(Lookup("error"), ShouldEqual, ERROR_MESSAGE)
		})

		Convey("'err' should map to ERROR_MESSAGE", func() {
			So(Lookup("err"), ShouldEqual, ERROR_MESSAGE)
		})

		Convey("'info' should map to INFO_MESSAGE", func() {
			So(Lookup("info"), ShouldEqual, INFO_MESSAGE)
		})

		Convey("'information' should map to INFO_MESSAGE", func() {
			So(Lookup("information"), ShouldEqual, INFO_MESSAGE)
		})

		Convey("'warning' should map to WARNING_MESSAGE", func() {
			So(Lookup("warning"), ShouldEqual, WARNING_MESSAGE)
		})

		Convey("'warn' should map to WARNING_MESSAGE", func() {
			So(Lookup("warn"), ShouldEqual, WARNING_MESSAGE)
		})

		Convey("'bob' should map to UNKNOWN", func() {
			So(Lookup("bob"), ShouldEqual, UNKNOWN)
		})
	})
}

func TestMessageTypeStringRepresentations(t *testing.T) {
	Convey("Each MessageType should have a string representation", t, func() {
		Convey("ERROR_MESSAGE should be represented by 'Error'", func() {
			So(ERROR_MESSAGE.String(), ShouldEqual, "Error")
		})

		Convey("INFO_MESSAGE should be represented by 'Information'", func() {
			So(INFO_MESSAGE.String(), ShouldEqual, "Information")
		})

		Convey("WARNING_MESSAGE should be represented by 'Warning'", func() {
			So(WARNING_MESSAGE.String(), ShouldEqual, "Warning")
		})

		Convey("UNKNOWN should be represented by 'Unknown'", func() {
			So(UNKNOWN.String(), ShouldEqual, "Unknown")
		})
	})
}
