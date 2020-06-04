package config

import (
	"testing"
)

func TestConfig_UnmarshalJson(t *testing.T) {
	// 本地文件系统中读取
	c := NewConfig(FromFile())
	var test1 interface{}
	if err := c.Unmarshal("test1.json", &test1); err != nil {
		t.Error(err)
	}

	if test1.(map[string]interface{})["int"].(float64) != 123 {
		t.Error(" int != 123 ")
	}

	if test1.(map[string]interface{})["string"].(string) != "abc" {
		t.Error(" string != abc ")
	}
}

func TestConfig_UnmarshalYaml(t *testing.T) {
	// 本地文件系统中读取
	c := NewConfig(FromFile())
	var test1 interface{}
	if err := c.Unmarshal("test1.yml", &test1); err != nil {
		t.Error(err)
	}

	if test1.(map[interface{}]interface{})["int"].(int) != 123 {
		t.Error(" int != 123 ")
	}

	if test1.(map[interface{}]interface{})["string"].(string) != "abc" {
		t.Error(" string != abc ")
	}
}