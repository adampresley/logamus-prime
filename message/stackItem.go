package message

import "strconv"

type StackItem struct {
	FileName     string
	LineNumber   int
	FunctionName string
}

func NewStackItem(fileName string, lineNumber int, functionName string) *StackItem {
	return &StackItem{
		FileName:     fileName,
		LineNumber:   lineNumber,
		FunctionName: functionName,
	}
}

func (this *StackItem) ToString() string {
	return "'" + this.FunctionName + "' on line " + strconv.Itoa(this.LineNumber) + " in " + this.FileName
}
