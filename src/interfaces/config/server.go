// Package config
package config

import (
	"fmt"

	"half-nothing.cn/service-core/interfaces/config"
)

type ServerConfig struct {
	HttpServerConfig *config.HttpServerConfig `yaml:"http"`
	GrpcServerConfig *config.GrpcServerConfig `yaml:"grpc"`
}

func (s *ServerConfig) InitDefaults() {
	s.HttpServerConfig = &config.HttpServerConfig{}
	s.HttpServerConfig.InitDefaults()
	s.GrpcServerConfig = &config.GrpcServerConfig{}
	s.GrpcServerConfig.InitDefaults()
}

func (s *ServerConfig) Verify() (bool, error) {
	if s.HttpServerConfig == nil {
		return false, fmt.Errorf("http server config is nil")
	}
	if ok, err := s.HttpServerConfig.Verify(); !ok {
		return false, err
	}
	if s.GrpcServerConfig == nil {
		return false, fmt.Errorf("grpc server config is nil")
	}
	if ok, err := s.GrpcServerConfig.Verify(); !ok {
		return false, err
	}
	return true, nil
}
