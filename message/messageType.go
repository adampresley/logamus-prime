package message

import (
	"strings"
)

type MessageType int

const (
	UNKNOWN         MessageType = 1
	ERROR_MESSAGE   MessageType = 2
	INFO_MESSAGE    MessageType = 3
	WARNING_MESSAGE MessageType = 4
)

var messageTypeStrings = map[MessageType]string{
	UNKNOWN:         "Unknown",
	ERROR_MESSAGE:   "Error",
	INFO_MESSAGE:    "Information",
	WARNING_MESSAGE: "Warning",
}

var messageTypeVariations = map[string]MessageType{
	"error":       ERROR_MESSAGE,
	"err":         ERROR_MESSAGE,
	"info":        INFO_MESSAGE,
	"information": INFO_MESSAGE,
	"warning":     WARNING_MESSAGE,
	"warn":        WARNING_MESSAGE,
}

func Lookup(messageType string) MessageType {
	messageType = strings.ToLower(messageType)

	for variation, possibleMessageType := range messageTypeVariations {
		if variation == messageType {
			return possibleMessageType
		}
	}

	return UNKNOWN
}

func (this MessageType) String() string {
	return messageTypeStrings[this]
}
