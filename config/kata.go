package config

type KataConfig struct {
	Username string
	Password string
	BaseURL  string
}

func GetKataConfig() *KataConfig {
	return &KataConfig{
		Username: GetString("KATA_USERNAME"),
		Password: GetString("KATA_PASSWORD"),
		BaseURL:  GetString("KATA_BASE_URL"),
	}
}
