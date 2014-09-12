#!/bin/bash

export LOGAMUS_SERVER_ADDRESS=localhost
export LOGAMUS_SERVER_PORT=9095

export LOGAMUS_SQL_ENGINE=mysql
export LOGAMUS_SQL_ADDRESS=localhost
export LOGAMUS_SQL_PORT=3306
export LOGAMUS_SQL_DATABASE=logamus
export LOGAMUS_SQL_USERNAME=root
export LOGAMUS_SQL_PASSWORD=password

go run logamus-prime.go
