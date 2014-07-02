package reader

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	t "time"

	"github.com/adampresley/logamus-prime/message"
)

type CFOutReader struct {
	scanner    *bufio.Scanner
	lineNumber int
}

func NewCFOutReader(reader io.Reader) (*CFOutReader, error) {
	result := &CFOutReader{
		scanner:    bufio.NewScanner(reader),
		lineNumber: 0,
	}

	return result, nil
}

func (this *CFOutReader) ReadMessage() (*message.Message, bool, error) {
	var date, time, messageBody string
	var messageType message.MessageType
	var err error

	/*
	 * Scan the next line from the file
	 */
	this.lineNumber++

	scanResult := this.scanner.Scan()

	if scanResult != true {
		return nil, scanResult, this.scanner.Err()
	}

	line := this.scanner.Text()

	/*
	 * Split the line by spaces. We are expecting ONE of the following structure.
	 *    - Date
	 *    - Time
	 *    - Thread
	 *    - Message Type
	 *    - Hyphen
	 *    - Message
	 *
	 * OR...
	 *
	 *    - Date
	 *    - Time
	 *    - Message Type
	 *    - Thread
	 *    - Message
	 *
	 * Note that the split is on a space character.
	 */
	parts := strings.Split(line, " ")

	if len(parts) < 5 {
		return nil, true, fmt.Errorf("CFOutReader: Line %d did not contain enough information to parse into a message", this.lineNumber)
	}

	// Date - It has no year, so put THIS year on it. Dumb
	now := t.Now()
	year := strconv.Itoa(now.Year())

	parsed, err := t.Parse("01/02/2006", parts[0]+"/"+year)
	if err != nil {
		date = parts[0]
	} else {
		date = parsed.Format("2006-01-02")
	}

	// Time
	time = parts[1]

	/*
	 * Thread or message type. If the string has brackets around
	 * it then we have a thread name.
	 */
	if strings.Contains(parts[2], "[") {
		messageType = message.Lookup(parts[3])
	} else {
		messageType = message.Lookup(parts[2])
	}

	/*
	 * We may have a hypen in position 4...
	 */
	if strings.Contains(parts[4], "-") {
		messageBody = strings.Join(parts[5:], " ")
	} else {
		messageBody = strings.Join(parts[4:], " ")
	}

	result := message.NewMessage(date, time, messageType, messageBody)

	/*
	 * Tag this message as "coldfusion" :)
	 */
	result.Tags = []string{"coldfusion", "log"}
	return result, true, nil
}
