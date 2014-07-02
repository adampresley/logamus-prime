/*
 * Script to create database tables for Logamus Prime.
 * This assumes the database has already been created.
 *
 * MySQL 5.5+
 */
CREATE TABLE messagetype (
	id INT PRIMARY KEY AUTO_INCREMENT,
	messageType VARCHAR(20)
);

CREATE TABLE message (
	id INT PRIMARY KEY AUTO_INCREMENT,
	dateTimeCreated DATE,
	messageType INT,
	message TEXT
);

CREATE TABLE messagestackitems (
	id INT PRIMARY KEY AUTO_INCREMENT,
	messageId INT,
	fileName VARCHAR(255),
	lineNumber INT,
	functionName VARCHAR(128)
);

CREATE TABLE messagetags (
	id INT PRIMARY KEY AUTO_INCREMENT,
	messageId INT,
	tag VARCHAR(50)
);

INSERT INTO messagetype (messageType) VALUES
	('Unknown'),
	('Error'),
	('Info'),
	('Warning')
;
