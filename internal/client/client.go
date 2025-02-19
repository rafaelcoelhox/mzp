package client

import (
	"context"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type whatsmeowClient struct {
	client *whatsmeow.Client
}

func NewClient(device *sqlstore.Device, logger waLog.Logger) WhatsAppClient {
	return &whatsmeowClient{
		client: whatsmeow.NewClient(device, logger),
	}
}

func (c *whatsmeowClient) Connect() error {
	return c.client.Connect()
}

func (c *whatsmeowClient) Disconnect() {
	c.client.Disconnect()
}

func (c *whatsmeowClient) AddEventHandler(handler interface{}) {
	c.client.AddEventHandler(handler)
}

func (c *whatsmeowClient) GetQRChannel(ctx context.Context) (<-chan whatsmeow.QRChannelItem, error) {
	return c.client.GetQRChannel(ctx)
}

func (c *whatsmeowClient) IsLoggedIn() bool {
	return c.client.Store.ID != nil
}
