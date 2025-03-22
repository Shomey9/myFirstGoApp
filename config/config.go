package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Feeds []string `json:"feeds"`
}

// LoadConfig reads and parses the configuration from a JSON file.
// It expects a command-line flag "config" to specify the path to the JSON file.
// If the flag is not provided, it defaults to "feeds.json".
// The function reads the file, unmarshals the JSON data into a Config struct,
// and returns a pointer to the Config struct or an error if any step fails.
//
// Usage:
//
//	go run main.go -config=path/to/your/config.json
//
// Returns:
//
//	*Config: A pointer to the Config struct populated with data from the JSON file.
//	error: An error if there is an issue reading the file or parsing the JSON.
func LoadConfig() (*Config, error) {
	configPath := flag.String("config", "feeds.json", "path to feeds JSON file")
	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	return &cfg, nil
}
