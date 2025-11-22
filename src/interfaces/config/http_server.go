// Package config
package config

import "fmt"

type HttpServerConfig struct {
	Enable bool   `yaml:"enable"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
}

func defaultHttpServerConfig() *HttpServerConfig {
	return &HttpServerConfig{
		Enable: true,
		Host:   "0.0.0.0",
		Port:   8080,
	}
}

func (h *HttpServerConfig) Verify() (bool, error) {
	if !h.Enable {
		return true, nil
	}
	if h.Host == "" {
		return false, fmt.Errorf("host is empty")
	}
	if h.Port <= 0 {
		return false, fmt.Errorf("port must larger than 0")
	}
	if h.Port > 65535 {
		return false, fmt.Errorf("port must less than 65535")
	}
	return true, nil
}
