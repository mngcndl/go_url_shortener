package config

type Config struct {
	StorageType string
	PostgresURL string
}

func LoadConfig(storageType, postgresURL string) *Config {
    return &Config{
        StorageType: storageType,
        PostgresURL: "postgres://user:password@localhost/dbname?sslmode=disable",
    }
}
