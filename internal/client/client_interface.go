package client

import (
	"context"

	"go.mau.fi/whatsmeow"
)

type WhatsAppClient interface {
	Connect() error
	Disconnect()
	AddEventHandler(handler interface{})
	GetQRChannel(ctx context.Context) (<-chan whatsmeow.QRChannelItem, error)
	IsLoggedIn() bool
}
