package sqlwriter

import (
	"database/sql"
	"errors"
	"time"

	"github.com/adampresley/logamus-prime/database"
	"github.com/adampresley/logamus-prime/message"
	"github.com/adampresley/logamus-prime/writer"
)

type SqlWriter struct {
	Db *sql.DB
}

func NewSqlWriter(writerType writer.WriterType, connectionInfo database.ConnectionInfo) (*SqlWriter, error) {
	var err error
	var result *SqlWriter
	var db *sql.DB

	switch writerType {
	case writer.MYSQL_WRITER:
		db, err = database.ConnectMySQL(connectionInfo)
		if err != nil {
			return nil, err
		}

	case writer.MSSQL_WRITER:
		db, err = database.ConnectMSSQL(connectionInfo)
		if err != nil {
			return nil, err
		}
	}

	result = &SqlWriter{Db: db}
	return result, nil
}

func (this *SqlWriter) Write(msg message.Message) error {
	/*
	 * Parse the date/time
	 */
	parsedDateTime, err := time.Parse("2006-01-02 15:04:05", msg.Date+" "+msg.Time)
	if err != nil {
		return errors.New("Invalid date/time format.")
	}

	transaction, err := this.Db.Begin()
	if err != nil {
		return errors.New("Error starting insert transaction: " + err.Error())
	}

	/*
	 * Insert the message and get the ID
	 */
	statement, err := transaction.Prepare("INSERT INTO message (dateTimeCreated, messageType, message) VALUES (?, ?, ?)")
	if err != nil {
		return errors.New("Error preparing insert statement for a new message: " + err.Error())
	}

	result, err := statement.Exec(
		parsedDateTime,
		msg.Type,
		msg.Message,
	)

	if err != nil {
		statement.Close()
		return errors.New("Error executing message insert: " + err.Error())
	}

	statement.Close()
	messageId, _ := result.LastInsertId()

	/*
	 * Insert any tags
	 */
	if len(msg.Tags) > 0 {
		for _, tag := range msg.Tags {
			statement, err = transaction.Prepare("INSERT INTO messagetags (messageId, tag) VALUES (?, ?)")
			if err != nil {
				transaction.Rollback()
				return errors.New("Unable to prepare tag insert: " + err.Error())
			}

			result, err = statement.Exec(
				int(messageId),
				tag,
			)

			if err != nil {
				statement.Close()
				transaction.Rollback()
				return errors.New("Unable to insert tag: " + err.Error())
			}

			statement.Close()
		}
	}

	/*
	 * Insert any stack trace items
	 */
	if len(msg.StackItems) > 0 {
		for _, stackItem := range msg.StackItems {
			statement, err = transaction.Prepare("INSERT INTO messagestackitems (messageId, fileName, lineNumber, functionName) VALUES (?, ?, ?, ?)")
			if err != nil {
				transaction.Rollback()
				return errors.New("Unable to prepare stack item insert: " + err.Error())
			}

			result, err = statement.Exec(
				int(messageId),
				stackItem.FileName,
				stackItem.LineNumber,
				stackItem.FunctionName,
			)

			if err != nil {
				statement.Close()
				transaction.Rollback()
				return errors.New("Unable to insert stack item: " + err.Error())
			}

			statement.Close()
		}
	}

	transaction.Commit()

	return nil
}
