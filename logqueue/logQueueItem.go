package logqueue

type LogQueueItem struct {
	LogType  LogReaderType
	FileName string
}

func NewFileLogQueueItem(logType LogReaderType, fileName string) *LogQueueItem {
	return &LogQueueItem{
		LogType:  logType,
		FileName: fileName,
	}
}
