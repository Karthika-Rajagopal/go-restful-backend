package config

// Config represents the application configuration
type Config struct {
	JWTSecret string
	// Add other configuration properties here
}

// LoadConfig loads the application configuration
func LoadConfig() *Config {
	return &Config{
		JWTSecret: "your-jwt-secret", 
		// Initialize other configuration properties here
	}
}
