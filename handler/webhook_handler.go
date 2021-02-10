package handler

import (
	"encoding/json"
	"example.com/m/contract"
	"example.com/m/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func PostWebhookHandler(dependencies service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postWebhookRequest, err := contract.NewPostWebhookRequest(r)

		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		let, err := dependencies.WebhookService.SendMessage(postWebhookRequest)

		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		response, err := json.Marshal(let)

		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	}
}
