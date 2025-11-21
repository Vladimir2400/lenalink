package sync

import (
	"fmt"
	"os"
	"time"
)

// Environment variable names for GARS configuration.
const (
	EnvGARSBaseURL  = "GARS_BASE_URL"
	EnvGARSUsername = "GARS_USERNAME"
	EnvGARSPassword = "GARS_PASSWORD"
	EnvGARSTimeout  = "GARS_TIMEOUT"
)

// Environment variable names for Aviasales configuration.
const (
	EnvAviasalesToken  = "AVIASALES_TOKEN"
	EnvAviasalesMarker = "AVIASALES_MARKER"
)

// Environment variable names for server configuration.
const (
	EnvListenAddr = "SYNC_LISTEN_ADDR"
)

// Default values for GARS.
const (
	DefaultGARSBaseURL  = "https://avibus.gars-ykt.ru:4443/avitest/odata/standard.odata"
	DefaultGARSUsername = "ХАКАТОН"
	DefaultGARSPassword = "123456"
	DefaultGARSTimeout  = 30 * time.Second
)

// Default values for server.
const (
	DefaultListenAddr = ":8080"
)

// Config contains configuration for all sync providers.
type Config struct {
	GARS      GARSConfig
	Aviasales AviasalesConfig
	RZD       RZDConfig
}

// GARSConfig contains configuration for connecting to GARS API.
type GARSConfig struct {
	BaseURL  string
	Username string
	Password string
	Timeout  time.Duration
}

// AviasalesConfig contains configuration for Aviasales API.
type AviasalesConfig struct {
	Token  string
	Marker string
}

// RZDConfig contains configuration for RZD API (currently mock).
type RZDConfig struct {
	// Future: API credentials will be added here
	Enabled bool
}

// LoadConfig reads complete sync configuration from environment.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		GARS:      LoadGARSConfig(),
		Aviasales: LoadAviasalesConfig(),
		RZD:       LoadRZDConfig(),
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}

// LoadGARSConfig reads GARS configuration from environment.
func LoadGARSConfig() GARSConfig {
	cfg := GARSConfig{
		BaseURL:  getEnvOrDefault(EnvGARSBaseURL, DefaultGARSBaseURL),
		Username: getEnvOrDefault(EnvGARSUsername, DefaultGARSUsername),
		Password: getEnvOrDefault(EnvGARSPassword, DefaultGARSPassword),
		Timeout:  DefaultGARSTimeout,
	}

	if raw := os.Getenv(EnvGARSTimeout); raw != "" {
		if d, err := time.ParseDuration(raw); err == nil {
			cfg.Timeout = d
		}
	}

	return cfg
}

// LoadAviasalesConfig reads Aviasales configuration from environment.
func LoadAviasalesConfig() AviasalesConfig {
	return AviasalesConfig{
		Token:  os.Getenv(EnvAviasalesToken),
		Marker: os.Getenv(EnvAviasalesMarker),
	}
}

// LoadRZDConfig reads RZD configuration from environment.
func LoadRZDConfig() RZDConfig {
	return RZDConfig{
		Enabled: true, // Always enabled for mock data
	}
}

// Validate ensures configuration is valid.
func (c *Config) Validate() error {
	if err := c.GARS.Validate(); err != nil {
		return fmt.Errorf("gars config: %w", err)
	}
	// Aviasales and RZD validation can be added here
	return nil
}

// Validate ensures GARS configuration is valid.
func (c GARSConfig) Validate() error {
	if c.BaseURL == "" {
		return fmt.Errorf("base url is required")
	}
	return nil
}

// getEnvOrDefault returns environment variable value or default.
func getEnvOrDefault(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
