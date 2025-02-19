package auth

import (
	"context"
	"fmt"
)

type QRAuthenticator struct {
	client WhatsAppClient
}

func NewQRAuthenticator(client WhatsAppClient) *QRAuthenticator {
	return &QRAuthenticator{client: client}
}

func (a *QRAuthenticator) Authenticate(ctx context.Context) error {
	if a.client.IsLoggedIn() {
		return a.client.Connect()
	}

	qrChan, err := a.client.GetQRChannel(ctx)
	if err != nil {
		return err
	}

	if err := a.client.Connect(); err != nil {
		return err
	}

	for evt := range qrChan {
		if evt.Event == "code" {
			fmt.Println("QR code:", evt.Code)
		} else {
			fmt.Println("Login event:", evt.Event)
		}
	}
	return nil
}
