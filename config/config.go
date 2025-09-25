package config

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func GetDbConfig() DbConfig {
	return DbConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "admin123",
		DbName:   "bioskop_db",
	}
}