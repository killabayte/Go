package storage

type Config struct {
	//Connection string to database
	DataBaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
