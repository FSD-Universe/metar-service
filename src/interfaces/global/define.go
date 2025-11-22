// Package global
package global

import (
	"context"
	"os"
	"strconv"
)

type Callable interface {
	Invoke(ctx context.Context) error
}

func CheckBoolEnv(envKey string, target *bool) {
	value := os.Getenv(envKey)
	if val, err := strconv.ParseBool(value); err == nil && val {
		*target = true
	}
}

func CheckStringEnv(envKey string, target *string) {
	value := os.Getenv(envKey)
	if value != "" {
		*target = value
	}
}
