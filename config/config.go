package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

// pgsql://yomeva:yomeva@db/consolidate
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "pgsql",
			Host:     "localhost",
			Port:     5432,
			Username: "yomeva",
			Password: "yomeva",
			Name:     "consolidate",
			Charset:  "utf8",
		},
	}
}
