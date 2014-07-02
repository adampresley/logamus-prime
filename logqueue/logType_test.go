package logqueue

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogReaderTypeStringRepresentations(t *testing.T) {
	Convey("Each LogReaderType should have a string representation", t, func() {
		Convey("CFOUT_LOG_TYPE should be represented by 'ColdFusion OUT File'", func() {
			So(CFOUT_LOG_TYPE.String(), ShouldEqual, "ColdFusion OUT File")
		})
	})
}
