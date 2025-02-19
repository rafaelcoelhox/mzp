package service

import "go.mau.fi/whatsmeow"

type WhatsAppRepository interface {
	Connect() error
}

type WhatsAppService struct {
	client *whatsmeow.Client
	repo   WhatsAppRepository
}

func NewWhatsAppService(repo WhatsAppRepository, client *whatsmeow.Client) *WhatsAppService {
	return &WhatsAppService{
		repo:   repo,
		client: client,
	}
}
