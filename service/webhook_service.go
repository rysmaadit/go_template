package service

import (
	"context"
	"errors"
	"example.com/m/contract"
	"example.com/m/external/kata"
	log "github.com/sirupsen/logrus"
)

type webhookService struct {
	kataClient kata.Client
}

type WebhookServiceInterface interface {
	SendMessage(req *contract.PostWebhookContract) (*kata.SendMessageResponse, error)
}

func NewWebhookService(kc kata.Client) *webhookService {
	return &webhookService{kataClient: kc}
}

func (ws *webhookService) SendMessage(req *contract.PostWebhookContract) (*kata.SendMessageResponse, error) {
	verifyContactContract := kata.VerifyContactContract{
		Blocking: kata.VerifyContactWait,
		Contacts: []string{req.RecipientNumber},
	}

	verifyContactResponse, err := ws.kataClient.VerifyContact(context.Background(), verifyContactContract)

	if err != nil {
		log.Error("error verify contact request", err)
		return nil, err
	}

	if verifyContactResponse.Contacts[0].Status != kata.ValidContactStatus {
		err = errors.New("contact not valid")
		log.Error("error verifying contact", err)
		return nil, err
	}

	localizableParams := []kata.LocalizableParam{
		{
			"default": req.RecipientName,
		},
	}

	language := kata.Language{
		Policy: kata.DeterministicPolicy,
		Code:   kata.LanguageID,
	}

	hsm := kata.HSM{
		Namespace:         kata.Namespace,
		ElementName:       kata.ElementName,
		Language:          language,
		LocalizableParams: localizableParams,
	}

	sendMessageTemplate := kata.SendMessageWithTemplateContract{
		To:            verifyContactResponse.Contacts[0].WAID,
		RecipientName: req.RecipientName,
		TTL:           kata.MessageTTL,
		Type:          kata.TypeHSM,
		HSM:           hsm,
	}

	return ws.kataClient.SendMessageWithTemplate(context.Background(), sendMessageTemplate)
}
