package config

import "os"

type Config struct {
	ServerPort string
	User                 string            // Username
	Passwd               string            // Password (requires User)
	Net                  string            // Network (e.g. "tcp", "tcp6", "unix". default: "tcp")
	Addr                 string            // Address (default: "127.0.0.1:3306" for "tcp" and "/tmp/mysql.sock" for "unix")
	DBName               string  
}

func InitConfig() *Config{
	return &Config{
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		User: getEnv("USER_DB", "root"),
		Passwd: getEnv("PASS_DB", "root"),
		Addr: getEnv("localhost", "127.0.0.1:3306"),
		Net: getEnv("NET", "tcp"),
		DBName: getEnv("DB_NAME", "todo-db"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}