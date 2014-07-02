package listener

import (
	"github.com/adampresley/logamus-prime/message"
	"github.com/adampresley/logamus-prime/writer"
)

func StartMessageListener(bufferSize int, writers []writer.MessageWriter, handler func([]writer.MessageWriter, message.Message)) chan message.Message {
	messageChannel := make(chan message.Message, bufferSize)

	go func() {
		for {
			message := <-messageChannel
			handler(writers, message)
		}
	}()

	return messageChannel
}
