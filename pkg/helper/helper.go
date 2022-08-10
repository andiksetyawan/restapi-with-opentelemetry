package helper

import (
	"os"
	"strconv"
)

//GetEnv with fallback value
func GetEnv(key, fallback string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	return fallback
}

func UintToString(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

func StringToUint(s string) (uint, error) {
	parseUint64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(parseUint64), nil
}
