package config

import (
    "os"
    "strconv"
    "strings"
)

type Config struct {
    TwitterServiceURL string
    DiscordServiceURL string
    Service          string
    APIKey           string
    SubnetIDs        []int
    Threshold        float64
    Interval         int
}

func LoadConfig() Config {
    return Config{
        TwitterServiceURL: os.Getenv("TWITTER_SERVICE_URL"),
        DiscordServiceURL: os.Getenv("DISCORD_SERVICE_URL"),
        Service:          os.Getenv("SERVICE"),
        APIKey:           os.Getenv("API_KEY"),
        SubnetIDs:        getEnvIntSlice("SUBNET_IDS", []int{14}),
        Threshold:        getEnvFloat("PERCENT_THRESHOLD", 5.0),
        Interval:         getEnvInt("CHECK_INTERVAL_MINUTES", 5),
    }
}

func getEnvIntSlice(name string, defaultVal []int) []int {
    if val, ok := os.LookupEnv(name); ok {
        parts := strings.Split(val, ",")
        results := make([]int, 0, len(parts))
        for _, part := range parts {
            part = strings.TrimSpace(part)
            if v, err := strconv.Atoi(part); err == nil {
                results = append(results, v)
            }
        }
        if len(results) > 0 {
            return results
        }
    }
    return defaultVal
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