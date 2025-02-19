package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rafaelcoelhox/myzap/internal/auth"
	"github.com/rafaelcoelhox/myzap/internal/client"
	"github.com/rafaelcoelhox/myzap/internal/handlers"
	"github.com/rafaelcoelhox/myzap/internal/storage"
)

type Application struct {
	config *Config
}

func NewApplication(cfg *Config) *Application {
	return &Application{config: cfg}
}

func (a *Application) Run() error {
	container, err := storage.NewContainer(a.config.DBURI, a.config.DBLog)
	if err != nil {
		return err
	}

	device, err := storage.GetDevice(container)
	if err != nil {
		return err
	}

	whatsAppClient := client.NewClient(device, a.config.ClientLog)
	messageHandler := handlers.NewMessageHandler()
	qrAuth := auth.NewQRAuthenticator(whatsAppClient)

	whatsAppClient.AddEventHandler(messageHandler.HandleEvent)

	ctx := context.Background()
	if err := qrAuth.Authenticate(ctx); err != nil {
		return err
	}

	a.waitForShutdown(whatsAppClient)
	return nil
}

func (a *Application) waitForShutdown(client client.WhatsAppClient) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	client.Disconnect()
}
