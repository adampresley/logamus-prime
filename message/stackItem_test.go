package message

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_StackItems(t *testing.T) {
	Convey("A StackItem object", t, func() {
		Convey("is returned by calling NewStackItem", func() {
			expected := &StackItem{
				FileName:     "test.go",
				LineNumber:   100,
				FunctionName: "runTest",
			}

			actual := NewStackItem("test.go", 100, "runTest")
			So(actual, ShouldResemble, expected)
		})

		Convey("converts to a pretty string", func() {
			stackItem := &StackItem{
				FileName:     "test.go",
				LineNumber:   100,
				FunctionName: "runTest",
			}

			expected := "'runTest' on line 100 in test.go"
			actual := stackItem.ToString()
			So(actual, ShouldEqual, expected)
		})
	})
}
