package config

import (
	"log" // to print log

	"github.com/BurntSushi/toml" // to read config.toml
)

// Represent database server n credentials
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil { // decode file config.toml
		log.Fatal(err)
	}
}
