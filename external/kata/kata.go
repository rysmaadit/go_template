package kata

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"example.com/m/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	VerifyContactWait   = "wait"
	ValidContactStatus  = "valid"
	DeterministicPolicy = "deterministic"
	LanguageID          = "id"
	Namespace           = "8eef29d7_fe14_49b3_9450_8758ce9cacb1"
	ElementName         = "lemonilo_official_store_greeting_2"
	MessageTTL          = "86400"
	TypeHSM             = "hsm"
)

type kataClient struct {
	config *config.KataConfig
	client *http.Client
}

type Client interface {
	VerifyContact(ctx context.Context, req VerifyContactContract) (*VerifyContactResponse, error)
	SendMessageWithTemplate(ctx context.Context, req SendMessageWithTemplateContract) (*SendMessageResponse, error)
}

func NewKataClient(config *config.KataConfig, client *http.Client) *kataClient {
	return &kataClient{
		config: config,
		client: client,
	}
}

func (kc *kataClient) VerifyContact(_ context.Context, req VerifyContactContract) (*VerifyContactResponse, error) {
	urlString := fmt.Sprintf("%s/contacts", kc.config.BaseURL)
	body, err := json.Marshal(req)

	if err != nil {
		log.Error("error marshal contact", err)
		return nil, err
	}

	b := bytes.NewBuffer(body)

	var res VerifyContactResponse

	token, err := kc.getToken()

	if err != nil {
		log.Error("error get token", err)
		return nil, err
	}

	err = kc.makeRequest(urlString, http.MethodPost, token, b, &res)

	if err != nil {
		log.Error("error validate contact", err)
		return nil, err
	}

	return &res, nil
}

func (kc *kataClient) SendMessageWithTemplate(_ context.Context, req SendMessageWithTemplateContract) (*SendMessageResponse, error) {
	urlString := fmt.Sprintf("%s/messages", kc.config.BaseURL)
	body, err := json.Marshal(req)

	if err != nil {
		log.Error("error marshal message request", err)
		return nil, err
	}

	b := bytes.NewBuffer(body)

	var res SendMessageResponse

	token, err := kc.getToken()

	if err != nil {
		log.Error("error get token", err)
		return nil, err
	}

	err = kc.makeRequest(urlString, http.MethodPost, token, b, &res)

	if err != nil {
		log.Error("error sending message", err)
		return nil, err
	}

	return &res, nil
}

func (kc *kataClient) makeRequest(url string, method string, token string, body io.Reader, dest interface{}) error {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		log.Error("error initiate request", err)
		return err
	}

	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := kc.client.Do(req)

	if err != nil {
		log.Error("error creating request", err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > http.StatusAccepted {
		bodyBytes, err := ioutil.ReadAll(res.Body)

		log.Info("status code", res.StatusCode)

		if err != nil {
			log.Error("status code error", err)
			return err
		}

		log.Error(string(bodyBytes))
		return errors.New(string(bodyBytes))
	}

	err = json.NewDecoder(res.Body).Decode(dest)

	if err != nil {
		log.Error("error decode body", err)
		return err
	}

	return nil
}

func (kc *kataClient) getToken() (string, error) {
	urlString := fmt.Sprintf("%s/users/login", kc.config.BaseURL)

	body, err := json.Marshal(LoginContract{
		Username: kc.config.Username,
		Password: kc.config.Password,
	})

	if err != nil {
		log.Error(err)
		return "", err
	}

	b := bytes.NewBuffer(body)

	var loginResponse LoginResponse

	err = kc.makeRequest(urlString, http.MethodPost, "", b, &loginResponse)

	if err != nil {
		log.Error("error get token", err)
		return "", err
	}

	return loginResponse.AccessToken, nil
}
