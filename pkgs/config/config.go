package config

import (
	"strconv"

	"github.com/joho/godotenv"

	"os"
)

func Load() {
	godotenv.Load(".env")
}

func Config(key string) string {
	return os.Getenv(key)
}

func ConfigInt(key string) int {
	strval := Config(key)
	val, err := strconv.Atoi(strval)
	if err != nil {
		panic(err)
	}
	return val
}
