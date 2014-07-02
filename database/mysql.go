package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(connectionInfo ConnectionInfo) (*sql.DB, error) {
	/*
	 * Create the connection
	 */
	log.Println("Connecting to MySQL database")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?autocommit=true",
		connectionInfo.UserName,
		connectionInfo.Password,
		connectionInfo.Host,
		connectionInfo.Port,
		connectionInfo.Database,
	))

	if err != nil {
		return nil, err
	}

	temp := os.Getenv("max_connections")
	if temp == "" {
		temp = "151"
	}

	maxConnections, _ := strconv.Atoi(temp)

	db.SetMaxIdleConns(maxConnections)
	db.SetMaxOpenConns(maxConnections)

	return db, nil
}
