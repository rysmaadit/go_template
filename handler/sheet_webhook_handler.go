package handler

import (
	"example.com/m/contract"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func PostWebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postWebhookRequest, err := contract.NewPostWebhookRequest(r)

		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Info(postWebhookRequest)
		w.WriteHeader(http.StatusOK)
	}
}
