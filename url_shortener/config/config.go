package config

type Config struct {
	StorageType string
	PostgresURL string
}

func LoadConfig(storageType, postgresURL string) *Config {
    return &Config{
        StorageType: storageType,
           PostgresURL: "postgres://user:password@postgres:5432/dbname?sslmode=disable",
    }
}
