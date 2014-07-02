package reader

import (
	"strings"
	"testing"

	"github.com/adampresley/logamus-prime/message"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewReaderFactories(t *testing.T) {
	Convey("Test creating new CF Out Reader object", t, func() {
		Convey("NewCFOutReader returns a new logReader object", func() {
			actual, _ := NewCFOutReader(strings.NewReader("This is a test"))
			So(actual, ShouldHaveSameTypeAs, &CFOutReader{})
		})
	})
}

func TestCFOutReader_ReadMessage(t *testing.T) {
	Convey("ReadMessage() correctly reads log lines, and will return errors when lines are incorrect", t, func() {
		Convey("Valid log line will return a Message", func() {
			reader, _ := NewCFOutReader(strings.NewReader("01/01 01:01:01 [thread-1] INFO - This is an informational message"))

			expected := &message.Message{
				Date:    "01/01",
				Time:    "01:01:01",
				Type:    message.INFO_MESSAGE,
				Message: "This is an informational message",
			}

			actual, _, _ := reader.ReadMessage()
			So(actual, ShouldResemble, expected)
		})

		Convey("Empty log line will return false and no message", func() {
			reader, _ := NewCFOutReader(strings.NewReader(""))
			_, actual, _ := reader.ReadMessage()
			So(actual, ShouldEqual, false)
		})

		Convey("Partial log line will return an error message", func() {
			reader, _ := NewCFOutReader(strings.NewReader("Bob Test"))
			_, _, actual := reader.ReadMessage()
			So(actual.Error(), ShouldEqual, "CFOutReader: Line 1 did not contain enough information to parse into a message")
		})
	})
}
