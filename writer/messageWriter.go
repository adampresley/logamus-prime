package writer

import "github.com/adampresley/logamus-prime/message"

type MessageWriter interface {
	Write(message.Message) error
}
