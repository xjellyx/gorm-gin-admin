package utils

import (
	"os"
	"strings"
)

// PubGetEnvString
func PubGetEnvString(key string, defaultValue string) (ret string) {
	ret = os.Getenv(key)
	if len(ret) == 0 {
		ret = defaultValue
	}
	return
}

// PubGetEnvBool
func PubGetEnvBool(key string, defaultValue bool) (ret bool) {
	val := strings.ToLower(os.Getenv(key))
	if val == "true" {
		ret = true
	} else if val == "false" {
		ret = false
	} else {
		ret = defaultValue
	}
	return
}
