package store

type Config struct {
	DatabaseUrl string `json:"databaseUrl"`
}

func NewConfig() *Config {
	return &Config{}
}
