package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectMSSQL(connectionInfo ConnectionInfo) (*sql.DB, error) {
	/*
	 * Create the connection
	 */
	log.Println("Connecting to MSSQL database")

	db, err := sql.Open("mssql", fmt.Sprintf("Server=%s;Port=%d;User Id=%s;Password=%s;Database=%s",
		connectionInfo.Host,
		connectionInfo.Port,
		connectionInfo.UserName,
		connectionInfo.Password,
		connectionInfo.Database,
	))

	if err != nil {
		return nil, err
	}

	return db, nil
}
