// Package config
package config

import "fmt"

type ServerConfig struct {
	HttpServerConfig *HttpServerConfig `yaml:"http"`
	GrpcServerConfig *GrpcServerConfig `yaml:"grpc"`
}

func defaultServerConfig() *ServerConfig {
	return &ServerConfig{
		HttpServerConfig: defaultHttpServerConfig(),
		GrpcServerConfig: defaultGrpcServerConfig(),
	}
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
