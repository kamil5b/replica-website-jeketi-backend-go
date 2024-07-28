package util

import (
	"os"
	"replica-website-jeketi-backend-go/constant"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return constant.EnvDefaultValues[key]
	}
	return value
}
