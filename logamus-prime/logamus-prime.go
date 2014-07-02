package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/adampresley/logamus-prime/database"
	"github.com/adampresley/logamus-prime/listener"
	"github.com/adampresley/logamus-prime/logqueue"
	"github.com/adampresley/logamus-prime/message"
	"github.com/adampresley/logamus-prime/writer"
	"github.com/adampresley/logamus-prime/writer/sqlwriter"
)

const (
	MAX_MESSAGE_LISTENER_GOROUTINES int = 20
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var err error
	start := time.Now()

	/*
	 * Setup channel and handler for CTRL+C (SIGINT)
	 */
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		for _ = range done {
			log.Println("Shutting down...")

			duration := time.Since(start)
			log.Println("Server ran for", duration.String())

			os.Exit(0)
		}
	}()

	/*
	 * Setup a SQL writer
	 */
	connectionInfo := database.ConnectionInfo{
		Host:     "localhost",
		Port:     3306,
		Database: "logamus",
		UserName: "root",
		Password: "password",
	}

	sqlWriter, err := sqlwriter.NewSqlWriter(writer.MYSQL_WRITER, connectionInfo)
	if err != nil {
		log.Println("Unable to setup MySQL writer:", err)
		os.Exit(1)
	}

	/*
	 * Start a log file listener. This will watch a directory and gobble up files,
	 * parse them, process messages, then delete the files.
	 */
	go listener.StartLogFileListener(logqueue.CFOUT_LOG_TYPE, "/home/adam/temp/cflogs-temp/*.log", []writer.MessageWriter{sqlWriter})()

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
	listener.StartHttpListener(":9095", []chan message.Message{messageChannel})
}
