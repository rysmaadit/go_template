package kata

type LoginContract struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type VerifyContactContract struct {
	Blocking string   `json:"blocking"`
	Contacts []string `json:"contacts"`
}

type VerifyContactResponse struct {
	Contacts []DetailContactResponse `json:"contacts"`
}

type DetailContactResponse struct {
	Input  string `json:"input"`
	Status string `json:"status"`
	WAID   string `json:"wa_id"`
}

type SendMessageWithTemplateContract struct {
	To            string `json:"to"`
	RecipientName string `json:"recipient_name"`
	TTL           string `json:"ttl"`
	Type          string `json:"type"`
	HSM           HSM    `json:"hsm"`
}

type HSM struct {
	Namespace         string             `json:"namespace"`
	ElementName       string             `json:"element_name"`
	Language          Language           `json:"language"`
	LocalizableParams []LocalizableParam `json:"localizable_params"`
}

type Language struct {
	Policy string `json:"policy"`
	Code   string `json:"code"`
}

type LocalizableParam map[string]interface{}

type SendMessageResponse struct {
	Messages []MessageID `json:"messages"`
}

type MessageID struct {
	ID string `json:"id"`
}
