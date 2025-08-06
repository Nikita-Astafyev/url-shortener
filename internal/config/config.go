package config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		Host:     "postgres",
		Port:     "5432",
		User:     "user",
		Password: "password",
		DBName:   "url_shortener",
	}
}

type ServerConfig struct {
	Port string
}

func GetServerConfig() ServerConfig {
	return ServerConfig{Port: "8080"}
}
