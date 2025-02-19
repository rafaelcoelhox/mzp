package storage

import (
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func NewContainer(dbURI string, logger waLog.Logger) (*sqlstore.Container, error) {
	return sqlstore.New("sqlite3", dbURI, logger)
}

func GetDevice(container *sqlstore.Container) (*sqlstore.Device, error) {
	return container.GetFirstDevice()
}
