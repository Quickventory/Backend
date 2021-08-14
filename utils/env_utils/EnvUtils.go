package env_utils

import (
	"fmt"
	"os"
)

func FetchEnvOrPanic(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("%s env variable is not set", key))
	}
	return value
}
