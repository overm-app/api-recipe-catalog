package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Env    string `yaml:"env"`
	Server struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
	Database struct {
		DBName string `yaml:"db_name"`
		Pool   struct {
			Max uint64 `yaml:"max"`
			Min uint64 `yaml:"min"`
		} `yaml:"pool"`
	} `yaml:"database"`

	MongoURI  string `yaml:"-"`
	JWTSecret string `yaml:"-"`
}

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config %s: %w", path, err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)
	var cfg Config
	if err := dec.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parse config %s: %w", path, err)
	}

	cfg.MongoURI = os.Getenv("MONGO_URI")
	cfg.JWTSecret = os.Getenv("AUTH_JWT_SECRET")

	if err := cfg.validate(); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) validate() error {
	var errs []error
	intPort, err := strconv.Atoi(c.Server.Port)
	if err != nil {
		errs = append(errs, err)
	}
	if intPort < 1 || intPort > 65535 {
		errs = append(errs, errors.New("server.port must be 1-65535"))
	}
	if c.Database.DBName == "" {
		errs = append(errs, errors.New("database.db_name is required"))
	}
	if c.MongoURI == "" {
		errs = append(errs, errors.New("MONGO_URI env is required"))
	}
	if len(c.JWTSecret) < 32 {
		errs = append(errs, errors.New("AUTH_JWT_SECRET env must be at least 32 chars"))
	}
	return errors.Join(errs...)
}
