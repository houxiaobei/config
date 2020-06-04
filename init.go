package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"path"
)

var cfg *Config

func Init(opts ...Option) {
	cfg = NewConfig(opts...)
}

func NewConfig(opts ...Option) *Config {
	c := &Config{}
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

//func GetContent(filename string) ([]byte, error) {
//	return cfg.GetContent(filename)
//}

func Unmarshal(filename string, v interface{}) error {
	content, err := cfg.GetContent(filename)
	if err != nil {
		return err
	}
	ext := path.Ext(filename)
	switch ext {
	case ".json":
		return json.Unmarshal(content, v)
	case ".yml":
		fallthrough
	case ".yaml":
		return yaml.Unmarshal(content, v)
	default:
		return &FileExtError{Ext: ext}
	}
}
