package message

type Message struct {
	Date       string       `json:"date"`
	Time       string       `json:"time"`
	Type       MessageType  `json:"type"`
	Message    string       `json:"message"`
	StackItems []*StackItem `json:"stackItems"`
	Tags       []string     `json:"tags"`
}

func NewMessage(date, time string, messageType MessageType, message string) *Message {
	return &Message{
		Date:    date,
		Time:    time,
		Type:    messageType,
		Message: message,
	}
}
