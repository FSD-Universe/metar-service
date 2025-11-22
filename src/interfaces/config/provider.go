// Package config
package config

import (
	"fmt"
	"slices"
	"strings"
)

type ProviderConfig struct {
	Type      string `yaml:"type"`
	Name      string `yaml:"name"`
	Target    string `yaml:"target"`
	Parser    string `yaml:"parser"`
	Selector  string `yaml:"selector"`
	Reverse   bool   `yaml:"reverse"`
	Multiline string `yaml:"multiline"`
}

var types = []string{"metar"}
var parsers = []string{"raw", "html", "json"}

func (p *ProviderConfig) Verify() (bool, error) {
	if p.Type == "" {
		return false, fmt.Errorf("type is required")
	}
	providerType := strings.ToLower(p.Type)
	if !slices.Contains(types, providerType) {
		return false, fmt.Errorf("type is not supported, only [%s] are supported", strings.Join(types, ","))
	}
	switch providerType {
	case "metar":
		if p.Name == "" {
			return false, fmt.Errorf("name is required")
		}
		if p.Target == "" {
			return false, fmt.Errorf("target is required")
		}
		if p.Parser == "" {
			return false, fmt.Errorf("parser is required")
		}
		parser := strings.ToLower(p.Parser)
		if !slices.Contains(parsers, parser) {
			return false, fmt.Errorf("parser is not supported, only [%s] are supported", strings.Join(parsers, ","))
		}
		switch parser {
		case "html":
			if p.Selector == "" {
				return false, fmt.Errorf("provider %s with type 'html' need a selector", p.Name)
			}
		case "json":
			if p.Selector == "" {
				return false, fmt.Errorf("provider %s with type 'json' need a selector", p.Name)
			}
		}
	}
	return true, nil
}
