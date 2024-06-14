package utils

import (
	"log"
	"os"
	"strconv"
)

func GetEnv(env string) string {
	value, ok := os.LookupEnv(env)
	if !ok {
		log.Fatalf("Environment variable %s not found", env)
	}
	return value
}

func GetIntEnv(env string) int {
	valueStr, ok := os.LookupEnv(env)
	if !ok {
		log.Fatalf("Environment variable %s not found", env)
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Invalid value for %s: %s", env, valueStr)
	}
	return value
}
