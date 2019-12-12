package config

// Log defines the available logging configuration.
type Log struct {
	Level  string
	Pretty bool
	Color  bool
}

type Config struct {
	File string
	Log  Log
}

// New initializes a new configuration with or without defaults.
func New() *Config {
	return &Config{}
}
