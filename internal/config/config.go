package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config represents the application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logger   LoggerConfig
}

// ServerConfig represents HTTP server configuration
type ServerConfig struct {
	Host            string
	Port            int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Driver          string // postgres, sqlite, memory
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	SSLMode         string
	MaxConnections  int
	MinConnections  int
	ConnectionTTL   time.Duration
	QueryTimeout    time.Duration
	MigrationsPath  string
}

// LoggerConfig represents logger configuration
type LoggerConfig struct {
	Level      string // DEBUG, INFO, WARN, ERROR
	JSONFormat bool
}

// Load loads configuration from environment variables and defaults
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Host:            getEnv("SERVER_HOST", "localhost"),
			Port:            getEnvInt("SERVER_PORT", 8080),
			ReadTimeout:     getEnvDuration("SERVER_READ_TIMEOUT", 15*time.Second),
			WriteTimeout:    getEnvDuration("SERVER_WRITE_TIMEOUT", 15*time.Second),
			ShutdownTimeout: getEnvDuration("SERVER_SHUTDOWN_TIMEOUT", 30*time.Second),
		},
		Database: DatabaseConfig{
			Driver:         getEnv("DB_DRIVER", "memory"),
			Host:           getEnv("DB_HOST", "localhost"),
			Port:           getEnvInt("DB_PORT", 5432),
			User:           getEnv("DB_USER", "postgres"),
			Password:       getEnv("DB_PASSWORD", ""),
			Database:       getEnv("DB_NAME", "lenalink"),
			SSLMode:        getEnv("DB_SSLMODE", "disable"),
			MaxConnections: getEnvInt("DB_MAX_CONNECTIONS", 25),
			MinConnections: getEnvInt("DB_MIN_CONNECTIONS", 5),
			ConnectionTTL:  getEnvDuration("DB_CONNECTION_TTL", 5*time.Minute),
			QueryTimeout:   getEnvDuration("DB_QUERY_TIMEOUT", 30*time.Second),
			MigrationsPath: getEnv("DB_MIGRATIONS_PATH", "./migrations"),
		},
		Logger: LoggerConfig{
			Level:      getEnv("LOG_LEVEL", "INFO"),
			JSONFormat: getEnvBool("LOG_JSON_FORMAT", false),
		},
	}
}

// Helper functions

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	val := os.Getenv(key)
	if val != "" {
		return val == "true" || val == "1" || val == "yes"
	}
	return defaultVal
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	valStr := os.Getenv(key)
	if valStr != "" {
		if duration, err := time.ParseDuration(valStr); err == nil {
			return duration
		}
	}
	return defaultVal
}

// ConnectionString returns the database connection string
func (dc DatabaseConfig) ConnectionString() string {
	switch dc.Driver {
	case "postgres":
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			dc.User, dc.Password, dc.Host, dc.Port, dc.Database, dc.SSLMode)
	case "sqlite":
		return dc.Database // SQLite expects a file path
	default:
		return ""
	}
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", c.Server.Port)
	}

	if c.Database.Driver == "" {
		return fmt.Errorf("database driver is required")
	}

	if c.Database.Driver == "postgres" {
		if c.Database.Host == "" {
			return fmt.Errorf("database host is required for postgres driver")
		}
		if c.Database.User == "" {
			return fmt.Errorf("database user is required for postgres driver")
		}
		if c.Database.Database == "" {
			return fmt.Errorf("database name is required for postgres driver")
		}
	}

	return nil
}
