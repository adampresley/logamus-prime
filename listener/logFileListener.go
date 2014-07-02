package listener

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/adampresley/logamus-prime/logqueue"
	"github.com/adampresley/logamus-prime/message"
	"github.com/adampresley/logamus-prime/reader"
	"github.com/adampresley/logamus-prime/writer"
)

func StartLogFileListener(logType logqueue.LogReaderType, globPattern string, writers []writer.MessageWriter) func() {
	return func() {
		for {
			filePaths, _ := filepath.Glob(globPattern)
			logItems := make([]*logqueue.LogQueueItem, 0)

			for _, filePath := range filePaths {
				logItems = append(logItems, logqueue.NewFileLogQueueItem(logType, filePath))
			}

			for _, logItem := range logItems {
				start := time.Now()

				var msg *message.Message
				var r reader.LogReader
				var fp *os.File
				var success bool
				var err error

				switch logType {
				case logqueue.CFOUT_LOG_TYPE:
					fp, err = os.Open(logItem.FileName)
					if err == nil {
						r, err = reader.NewCFOutReader(bufio.NewReader(fp))
					}
					defer fp.Close()
				}

				/*
				 * If the file opened successfully...
				 */
				if err == nil {
					success = true

					/*
					 * Parse lines until we are done
					 */
					for success {
						msg, success, err = r.ReadMessage()

						/*
						 * If a line was read, but there was an error, log it.
						 * Otherwise it we read a line, sent it to listening writers
						 */
						if success != true && err != nil {
							log.Println("ERROR - Error in reading message:", err)
						} else if success == true {
							if msg != nil {
								go func(msg message.Message) {
									for _, w := range writers {
										w.Write(msg)
									}
								}(*msg)
							}
						}
					}

					os.Remove(logItem.FileName)
				} else {
					log.Println("ERROR - Could not create reader...", err)
				}

				duration := time.Since(start)
				log.Println("Parsed", logItem.FileName, "in", duration.String())

			}
		}
	}
}
