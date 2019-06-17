package config

import (
	"os"
	"strconv"
)

type myDBConfig struct {
	IP       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Timeout  int    `json:"timeout"`
	Name     string `json:"name"`
}

func GetDBConfig() myDBConfig {
	config := myDBConfig{
		IP:       getEnvString("DB_IP", "0.0.0.0"),
		Username: getEnvString("DB_USERNAME", ""),
		Password: getEnvString("DB_PASSWORD", ""),
		Port:     getEnvInt("DB_PORT", 0),
		Timeout:  getEnvInt("DB_TIMEOUT", 0),
		Name:     getEnvString("DB_NAME", ""),
	}
	return config
}

func getEnvInt(name string, defaultVal int) int {
	valueStr := os.Getenv(name)
	if valueStr != "" {
		if valueInt, err := strconv.Atoi(valueStr); err == nil {
			return valueInt
		}
	}
	return defaultVal
}

func getEnvString(name string, defaultVal string) string {
	valueStr := os.Getenv(name)
	if valueStr != "" {
		return valueStr
	}
	return defaultVal
}
