package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server *Server      `yaml:"server,omitempty"`
	App    *Application `yaml:"application,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

// Application holds application configuration details
type Application struct {
	DefaultDbhost          string `yaml:"default_db_host"`
	DefaultElasticBulkType string `yaml:"default_elastic_bulk_type"`
	DefaultElasticHost     string `yaml:"default_elastic_host"`
	DefaultFetch           int    `yaml:"default_fetch"`
	DefaultOffset          int    `yaml:"default_offset"`
	MinPasswordStr         int    `yaml:"min_password_strength,omitempty"`
	SwaggerUIPath          string `yaml:"swagger_ui_path,omitempty"`
}
