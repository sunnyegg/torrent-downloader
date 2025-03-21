package config

import (
	"github.com/anacrolix/torrent"
)

// Config holds the application configuration
type Config struct {
	DownloadDir string
	Client      *torrent.ClientConfig
}

// NewDefaultConfig creates a new configuration with default settings
func NewDefaultConfig() *Config {
	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = "./downloads"

	return &Config{
		DownloadDir: "./downloads",
		Client:      clientConfig,
	}
}
