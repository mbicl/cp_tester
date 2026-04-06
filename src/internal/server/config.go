package server

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"go.yaml.in/yaml/v2"
)

type Config struct {
	CPPath string `yaml:"CPPATH"`
}

func LoadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// Check if the config file exists, if not, create a default config file.
	if err := checkConfigPath(configPath); err != nil {
		return nil, err
	}

	// Load the config file and parse it.
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func getConfigPath() (string, error) {
	var configPath string
	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "", os.ErrNotExist
		}
		configPath = filepath.Join(appData, "cp_tester", "config.yaml")
	case "darwin", "linux":
		configPath = filepath.Join(os.Getenv("HOME"), ".config", "cp_tester", "config.yaml")
	default:
		return "", errors.New("unsupported platform")
	}
	return configPath, nil
}

func checkConfigPath(configPath string) error {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Printf("Config file not found at %s. Creating a default config file.", configPath)
		e := os.MkdirAll(filepath.Dir(configPath), 0755)
		if e != nil {
			return e
		}
		e = os.WriteFile(configPath, []byte("CPPATH: ~/cp/\n"), 0644)
		if e != nil {
			return e
		}
		return nil
	}
	return err
}
