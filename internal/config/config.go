package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	data map[string]map[string]string
}

func Load(path string) (*Config, error) {

	fmt.Println("Loading config:", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("CFG001: configuration file not found")
	}
	defer file.Close()

	cfg := &Config{
		data: make(map[string]map[string]string),
	}

	section := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {

			section = strings.TrimSpace(line[1 : len(line)-1])

			if _, ok := cfg.data[section]; !ok {
				cfg.data[section] = make(map[string]string)
			}

			continue
		}

		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		cfg.data[section][key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Get(section, key string) string {

	if sec, ok := c.data[section]; ok {

		if value, ok := sec[key]; ok {
			return value
		}
	}

	return ""
}

// ---------- Helper ----------

func (c *Config) ServerPort() string {

	port := c.Get("server", "port")

	if port == "" {
		return "8001"
	}

	return port
}

func (c *Config) MaxVolumeSize() string {
	return c.Get("storage", "max_volume_size")
}

func (c *Config) Volumes() map[string]string {

	volumes := make(map[string]string)

	for section, values := range c.data {

		if strings.HasPrefix(section, "volume:") {

			name := strings.TrimPrefix(section, "volume:")

			volumes[name] = values["path"]
		}
	}

	return volumes
}

func (c *Config) Volume(name string) string {

	return c.Get("volume:"+name, "path")
}
