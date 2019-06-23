package config

import (
	"os"
	"strconv"
)

// MyDBConfig database config values
type MyDBConfig struct {
	IP       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Timeout  int    `json:"timeout"`
	Name     string `json:"name"`
}

// GetDBConfig returns the database configurations
func GetDBConfig() MyDBConfig {
	config := MyDBConfig{
		IP:       getEnvString("DB_IP", ""),
		Username: getEnvString("DB_USERNAME", ""),
		Password: getEnvString("DB_PASSWORD", ""),
		Port:     getEnvInt("DB_PORT", 0),
		Timeout:  getEnvInt("DB_TIMEOUT", 0),
		Name:     getEnvString("DB_NAME", ""),
	}
	return config
}

func getEnvInt(name string, defaultVal int) int {
	valueStr, _ := os.LookupEnv(name)
	if valueStr != "" {
		if valueInt, err := strconv.Atoi(valueStr); err == nil {
			return valueInt
		}
	}
	return defaultVal
}

func getEnvString(name string, defaultVal string) string {
	valueStr, _ := os.LookupEnv(name)
	if valueStr != "" {
		return valueStr
	}
	return defaultVal
}
