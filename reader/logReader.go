package reader

import (
	"github.com/adampresley/logamus-prime/message"
)

type LogReader interface {
	ReadMessage() (*message.Message, bool, error)
}
