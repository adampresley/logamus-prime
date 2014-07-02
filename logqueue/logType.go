package logqueue

type LogReaderType int

const (
	CFOUT_LOG_TYPE LogReaderType = iota
)

var logTypeStrings = map[LogReaderType]string{
	CFOUT_LOG_TYPE: "ColdFusion OUT File",
}

func (this LogReaderType) String() string {
	return logTypeStrings[this]
}
