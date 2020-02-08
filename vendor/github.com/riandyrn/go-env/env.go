package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetString is used for fetching string value from env.
// Returns empty string if key is not exist.
func GetString(key string) string {
	return os.Getenv(key)
}

// GetInt is used for fetching int value from env.
// Returns `0` if key is not exist or has invalid value.
func GetInt(key string) int {
	v, _ := strconv.Atoi(GetString(key))
	return v
}

// GetBool is used for fetching bool value from env.
// Returns `false` if key is not exist or has invalid value.
func GetBool(key string) bool {
	v, _ := strconv.ParseBool(GetString(key))
	return v
}

// GetSeconds is used for fetching duration in seconds from env.
// Returns `0` seconds if key is not exists.
func GetSeconds(key string) time.Duration {
	v := GetInt(key)
	return time.Duration(v) * time.Second
}

// GetStrings is used for fetching slice of string from env.
// Returns `nil` if key is not exists.
func GetStrings(key string, sep string) []string {
	v := GetString(key)
	return strings.Split(v, sep)
}
