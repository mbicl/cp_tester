package server

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"go.yaml.in/yaml/v2"
)

type Config struct {
	CPPath   string `yaml:"CPPATH"`
	Language string `yaml:"LANGUAGE"`
	Template string `yaml:"TEMPLATE"`
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
	config := loadDefaultConfig()
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	config.Language = strings.ToLower(config.Language)
	config.CPPath = expandHome(config.CPPath)
	config.Template = expandHome(config.Template)
	return &config, nil
}

func loadDefaultConfig() Config {
	return Config{
		CPPath:   "~/cp/",
		Language: "cpp",
		Template: "",
	}
}

func expandHome(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return filepath.Join(home, path[2:])
	}
	return path
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
		e = os.WriteFile(configPath, []byte("CPPATH: ~/cp/\nLANGUAGE: cpp\n"), 0644)
		if e != nil {
			return e
		}
		return nil
	}
	return err
}
