package environment

import (
	"fmt"
	"os"
	"strconv"

	"github.com/adampresley/logamus-prime/database"
)

/*
Reads environment variables to get the server address
and port binding information.
*/
func GetServerBindingInformation() (string, int, error) {
	var err error

	address := os.Getenv("LOGAMUS_SERVER_ADDRESS")
	if address == "" {
		return "", 0, fmt.Errorf("Invalid LOGAMUS_SERVER_ADDRESS environment variable")
	}

	portString := os.Getenv("LOGAMUS_SERVER_PORT")
	if portString == "" {
		return "", 0, fmt.Errorf("Invlid LOGAMUS_SERVER_PORT environment variable")
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return "", 0, fmt.Errorf("Invalid Logamus port: %s", err.Error())
	}

	return address, port, nil
}

/*
Reads environment variables to get SQL server connection
information and returns a ConnectionInfo object. The following
constants represent the environment variables that are read:
*/
func GetSqlConnectionInformation() (string, database.ConnectionInfo, error) {
	var err error
	var connectionInfo database.ConnectionInfo

	engine := os.Getenv("LOGAMUS_SQL_ENGINE")
	if engine == "" {
		return "", connectionInfo, fmt.Errorf("Invalid ENGINE environment variable")
	}

	address := os.Getenv("LOGAMUS_SQL_ADDRESS")
	if address == "" {
		return "", connectionInfo, fmt.Errorf("Invalid environment variable LOGAMUS_SQL_ADDRESS")
	}

	portString := os.Getenv("LOGAMUS_SQL_PORT")
	if portString == "" {
		return "", connectionInfo, fmt.Errorf("Invalid environment variable LOGAMUS_SQL_PORT")
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return "", connectionInfo, fmt.Errorf("Invalid SQL port: %s", err.Error())
	}

	databaseName := os.Getenv("LOGAMUS_SQL_DATABASE")
	if databaseName == "" {
		return "", connectionInfo, fmt.Errorf("Invalid environment variable LOGAMUS_SQL_DATABASE")
	}

	userName := os.Getenv("LOGAMUS_SQL_USERNAME")
	if userName == "" {
		return "", connectionInfo, fmt.Errorf("Invalid environment variable LOGAMUS_SQL_USERNAME")
	}

	password := os.Getenv("LOGAMUS_SQL_PASSWORD")
	if password == "" {
		return "", connectionInfo, fmt.Errorf("Invalid environment variable LOGAMUS_SQL_PASSWORD")
	}

	connectionInfo = database.ConnectionInfo{
		Host:     address,
		Port:     port,
		Database: databaseName,
		UserName: userName,
		Password: password,
	}

	return engine, connectionInfo, nil
}