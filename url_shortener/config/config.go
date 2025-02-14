package config

// import "os"

type Config struct {
	StorageType string
	PostgresURL string
}

func LoadConfig() *Config {
	return &Config{
		StorageType: "postgres",
		PostgresURL: "postgres://user:password@localhost/dbname?sslmode=disable",
	}
}