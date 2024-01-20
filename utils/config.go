package utils

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Port                 uint16
	MOTD                 string
	OnlineMode           bool
	ViewDistance         int
	CompressionThreshold int
	VersionProtocol      int
	VersionName          string
	MaxPlayers           int
}

func LoadConfig() ServerConfig {
	config := ServerConfig{
		Port:                 25565,
		MOTD:                 "A Minecraft Server",
		OnlineMode:           false,
		ViewDistance:         10,
		CompressionThreshold: 256,
		VersionProtocol:      755,
		VersionName:          "1.20.4",
	}

	serverDirectory := ServerDirectory()

	if _, err := os.Stat(serverDirectory); os.IsNotExist(err) {
		os.Mkdir(serverDirectory, 0755)
	}

	if _, err := os.Stat(serverDirectory + "/server.yml"); err == nil {
		loadedConfig, err := ioutil.ReadFile(serverDirectory + "/server.yml")
		if err != nil {
			Error("Failed to load server.yml")
		}

		err = yaml.Unmarshal(loadedConfig, &config)
		if err != nil {
			Error("Failed to parse server.yml")
		}

		Info("Loaded server.yml")

		return config
	} else {
		Warn("server.yml not found, using default configuration (PATH: " + serverDirectory + "/server.yml)")

		configBytes, err := yaml.Marshal(config)
		if err != nil {
			Error("Failed to create server.yml")
			return config
		}

		err = ioutil.WriteFile(serverDirectory+"/server.yml", configBytes, 0644)

		if err != nil {
			Error("Failed to create server.yml")
		}

		Info("Created server.yml")
		return config
	}
}
