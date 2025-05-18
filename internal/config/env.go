package config

import (
    "os"
    "strconv"
)

type Config struct {
	TwitterServiceURL string
	APIKey string
    SubnetID  int
    Threshold float64
    Interval  int
}

func LoadConfig() Config {
    return Config{
		TwitterServiceURL: os.Getenv("TWITTER_SERVICE_URL"),
		APIKey: os.Getenv("API_KEY"),
        SubnetID:    getEnvInt("SUBNET_ID", 11),
        Threshold:   getEnvFloat("PERCENT_THRESHOLD", 5.0),
        Interval:    getEnvInt("CHECK_INTERVAL_MINUTES", 5),
    }
}

func getEnvInt(name string, defaultVal int) int {
    if val, ok := os.LookupEnv(name); ok {
        if v, err := strconv.Atoi(val); err == nil {
            return v
        }
    }
    return defaultVal
}

func getEnvFloat(name string, defaultVal float64) float64 {
    if val, ok := os.LookupEnv(name); ok {
        if v, err := strconv.ParseFloat(val, 64); err == nil {
            return v
        }
    }
    return defaultVal
}