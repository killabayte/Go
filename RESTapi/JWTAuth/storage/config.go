package storage

type Config struct {
	//Connection string to database
	DataBaseURL string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
