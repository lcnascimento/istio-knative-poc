package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetString ...
func GetString(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return value
}

// MustGetString ...
func MustGetString(name string) string {
	value := os.Getenv(name)
	if value == "" {
		fmt.Printf("%s can't be empty\n", name)
		os.Exit(1)
	}
	return value
}

// GetInt ...
func GetInt(name string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		value = defaultValue
	}
	return value
}

// MustGetInt ...
func MustGetInt(name string) int {
	value, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		fmt.Printf("%s must contain a integer value!\n", name)
		os.Exit(1)
	}
	return value
}

// GetFloat ...
func GetFloat(name string, defaultValue float64) float64 {
	value, err := strconv.ParseFloat(os.Getenv(name), 64)
	if err != nil {
		value = defaultValue
	}
	return value
}

// MustGetFloat ...
func MustGetFloat(name string) float64 {
	value, err := strconv.ParseFloat(os.Getenv(name), 64)
	if err != nil {
		fmt.Printf("%s must contain a float value!\n", name)
		os.Exit(1)
	}
	return value
}

// GetBool ...
func GetBool(name string, defaultValue bool) bool {
	value, err := strconv.ParseBool(os.Getenv(name))
	if err != nil {
		value = defaultValue
	}
	return value
}

// MustGetBool ...
func MustGetBool(name string) bool {
	value, err := strconv.ParseBool(os.Getenv(name))
	if err != nil {
		fmt.Printf("%s must contain a boolean value! (true or false)\n", name)
		os.Exit(1)
	}
	return value
}
