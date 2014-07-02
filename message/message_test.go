package message

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_NewMessage_ReturnsMessageObject(t *testing.T) {
	Convey("Calling NewMessage returns a pointer to a Message object", t, func() {
		expected := &Message{
			Date:    "01/01",
			Time:    "01:01:01",
			Type:    INFO_MESSAGE,
			Message: "Test message",
		}

		actual := NewMessage("01/01", "01:01:01", INFO_MESSAGE, "Test message")
		So(actual, ShouldResemble, expected)
	})
}
