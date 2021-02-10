package contract

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type PostWebhookContract struct {
	Date            string `json:"date"`
	SourceOrder     string `json:"source_order"`
	OrderID         string `json:"order_id"`
	InvoiceCode     string `json:"invoice_code"`
	CustomerName    string `json:"customer_name"`
	RecipientNumber string `json:"recipient_number"`
	RecipientName   string `json:"recipient_name"`
}

func NewPostWebhookRequest(r *http.Request) (*PostWebhookContract, error) {
	var postWebhookContract PostWebhookContract

	if err := json.NewDecoder(r.Body).Decode(&postWebhookContract); err != nil {
		log.Error(err)
		return nil, err
	}

	return &postWebhookContract, nil
}
