package logqueue

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewLogQueueItems(t *testing.T) {
	Convey("Test creating new LogQueueItem objects", t, func() {
		Convey("NewFileLogQueueItem will return a new LogQueueItem object with file name and log type populated", func() {
			expected := &LogQueueItem{
				LogType:  CFOUT_LOG_TYPE,
				FileName: "test.log",
			}

			actual := NewFileLogQueueItem(CFOUT_LOG_TYPE, "test.log")
			So(actual, ShouldResemble, expected)
		})
	})
}
