package config

type Config struct {
	PostgresDsn string `yaml:"postgres_dsn"`
	ServerPort  string `yaml:"server_port"`
}
