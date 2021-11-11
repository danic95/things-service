package settings

import (
	"context"
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type HeaderKey string

const (
	defaultFilePath = "settings.yml"
)

var ErrNoFile error = errors.New("file not found")
var ErrParsingFile error = errors.New("unable to parse file")

type Settings struct {
	Service Service  `yaml:"Service"`
	Cache   Cache    `yaml:"Cache"`
	DB      Database `yaml:"Database"`
}

type Service struct {
	Name       string `yaml:"name"`
	PathPrefix string `yaml:"path_prefix"`
	Version    string `yaml:"version"`
	Port       int    `yaml:"port"`
	Debug      bool   `yaml:"debug"`
}

type Cache struct {
	Enabled bool   `yaml:"enabled"`
	Addr    string `yaml:"addr"`
}

type Database struct {
	Engine   string `yaml:"engine"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// GetEnv returns env value, if empty, returns fallback value
func GetEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

//Get returns main configuration object
func Get(ctx context.Context, filePath string) (*Settings, error) {

	var err error
	var confFile []byte

	config := &Settings{}

	confFile, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, ErrNoFile
	}

	//if file exists use its variables

	err = yaml.Unmarshal(confFile, &config)
	if err != nil {
		return nil, ErrParsingFile
	}

	return config, nil
}

func New(ctx context.Context) (*Settings, error) {
	return Get(ctx, defaultFilePath)
}
