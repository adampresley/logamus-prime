package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/adampresley/logamus-prime/database"
	"github.com/adampresley/logamus-prime/environment"
	"github.com/adampresley/logamus-prime/listener"
	"github.com/adampresley/logamus-prime/message"
	"github.com/adampresley/logamus-prime/writer"
	"github.com/adampresley/logamus-prime/writer/sqlwriter"
	"github.com/adampresley/sigint"
)

const (
	MAX_MESSAGE_LISTENER_GOROUTINES int = 20
)

func main() {
	var err error

	runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()

	/*
	 * Setup channel and handler for CTRL+C (SIGINT)
	 */
	sigint.ListenForSIGINT(func() {
		log.Println("Shutting down...")

		duration := time.Since(start)
		log.Println("Server ran for", duration.String())

		os.Exit(0)
	})

	/*
	 * Get our server address binding information
	 */
	serverAddress, serverPort, err := environment.GetServerBindingInformation()
	if err != nil {
		log.Fatalf("There was an error getting server binding information: %s", err.Error())
	}

	/*
	 * Setup a SQL writer
	 */
	engine, connectionInfo, err := environment.GetSqlConnectionInformation()
	if err != nil {
		log.Fatalf("There was error getting SQL connection information: %s", err.Error())
	}

	sqlWriter := setupSqlWriter(engine, &connectionInfo)

	/*
	 * Setup a listener for messages that are recieved.
	 */
	messageChannel := listener.StartMessageListener(MAX_MESSAGE_LISTENER_GOROUTINES, []writer.MessageWriter{sqlWriter}, func(writers []writer.MessageWriter, msg message.Message) {
		var tags string

		log.Println("Type", msg.Type.String(), "@", msg.Date, msg.Time, "-", msg.Message)

		if len(msg.StackItems) > 0 {
			log.Println("Stack Trace:")

			for _, stackItem := range msg.StackItems {
				log.Println("Function call to", stackItem.FunctionName, "on line", stackItem.LineNumber, "in", stackItem.FileName)
			}
		}

		if len(msg.Tags) > 0 {
			tags = ""
			for _, tag := range msg.Tags {
				tags += tag + ", "
			}

			log.Println("Tags:", tags)
		}

		for _, w := range writers {
			w.Write(msg)
		}
	})

	/*
	 * Fire up the HTTP listener. POSTS to the HTTP server will
	 * send parsed messages to any message listeners setup.
	 */
	listener.StartHttpListener(fmt.Sprintf("%s:%d", serverAddress, serverPort), []chan message.Message{messageChannel})
}

func setupSqlWriter(writerName string, connectionInfo *database.ConnectionInfo) writer.MessageWriter {
	var result writer.MessageWriter
	var err error

	switch writerName {
	case "mysql":
		result, err = sqlwriter.NewSqlWriter(writer.MYSQL_WRITER, *connectionInfo)

	case "mssql":
		result, err = sqlwriter.NewSqlWriter(writer.MSSQL_WRITER, *connectionInfo)
	}

	if err != nil {
		log.Println("Unable to setup SQL writer:", err)
		os.Exit(1)
	}

	return result
}

