package app

import (
	waLog "go.mau.fi/whatsmeow/util/log"
)

type Config struct {
	DBURI     string
	DBLog     waLog.Logger
	ClientLog waLog.Logger
}

func NewConfig() *Config {
	return &Config{
		DBURI:     "file:tmpstore.db?_foreign_keys=on",
		DBLog:     waLog.Stdout("Database", "DEBUG", true),
		ClientLog: waLog.Stdout("Client", "DEBUG", true),
	}
}
