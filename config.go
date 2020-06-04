package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"path"
)

type Config struct {
	drivers []drivers
}

type drivers interface {
	Get(string) ([]byte, error)
}

//func (c *Config) apply(o Option) {
//	c.drivers = append(c.drivers,o.)
//}

func (c *Config) GetContent(filename string) ([]byte, error) {
	for _, driver := range c.drivers {
		content, err := driver.Get(filename)
		if err != nil {
			continue
		}
		return content, nil
	}
	return nil, &FileNotFound{File: filename}
}

func (c *Config) Unmarshal(filename string, v interface{}) error {
	content, err := c.GetContent(filename)
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
