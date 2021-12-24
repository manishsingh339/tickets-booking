package config

type DatabaseConfig struct {
	Name     string
	Host     string
	Port     string
	URL      string
	User     string
	Password string
}

var dbConfig DatabaseConfig

func GetDBConfig() DatabaseConfig {
	return DatabaseConfig{
		User:     "",
		Password: "",
		Name:     "ticket_booking",
		Host:     "localhost",
		URL:      "mongodb://localhost",
		Port:     "27017",
	}
}
