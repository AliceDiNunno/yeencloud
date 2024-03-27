package env

import (
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	cli *cli.Context
}

func NewConfig(c *cli.Context) *Config {
	return &Config{
		cli: c,
	}
}

func (config *Config) getEnvStringOrDefault(key string, defaultValue string) string {
	// CLI takes precedence over env variables.
	value := config.cli.String(key)

	if value != "" {
		return value
	}

	envVariable, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	if strings.HasPrefix(envVariable, "$") {
		varNameWithoutPrefix := strings.TrimPrefix(envVariable, "$")
		return config.getEnvStringOrDefault(varNameWithoutPrefix, defaultValue)
	}

	return envVariable
}

func (config *Config) GetEnvBoolOrDefault(key string, defaultValue bool) bool {
	value := config.cli.Bool(key)

	if value {
		return value
	}

	envVariable, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	if strings.HasPrefix(envVariable, "$") {
		varNameWithoutPrefix := strings.TrimPrefix(envVariable, "$")
		return config.GetEnvBoolOrDefault(varNameWithoutPrefix, defaultValue)
	}

	convertedValue, err := strconv.ParseBool(envVariable)

	if err != nil {
		return defaultValue
	}

	return convertedValue
}

func (config *Config) GetEnvStringOrDefault(key string, defaultValue string) string {
	return config.getEnvStringOrDefault(key, defaultValue)
}

func (config *Config) getEnvIntOrDefault(key string, defaultValue int) int {
	value := config.cli.Int(key)

	if value != 0 {
		return value
	}

	envVariable, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	if strings.HasPrefix(envVariable, "$") {
		varNameWithoutPrefix := strings.TrimPrefix(envVariable, "$")
		return config.getEnvIntOrDefault(varNameWithoutPrefix, defaultValue)
	}

	convertedValue, err := strconv.Atoi(envVariable)

	if err != nil {
		return defaultValue
	}

	return convertedValue
}

func (config *Config) GetEnvIntOrDefault(key string, defaultValue int) int {
	return config.getEnvIntOrDefault(key, defaultValue)
}
