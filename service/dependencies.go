package service

import (
	"example.com/m/config"
	"example.com/m/external/kata"
	"net/http"
)

type Dependencies struct {
	WebhookService WebhookServiceInterface
}

func InstantiateDependencies() Dependencies {
	kataClient := kata.NewKataClient(config.GetKataConfig(), &http.Client{})
	webhookService := NewWebhookService(kataClient)
	return Dependencies{
		WebhookService: webhookService,
	}
}
