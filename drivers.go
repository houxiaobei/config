package config

import "io/ioutil"

type FileDrive struct {
}

func (l *FileDrive) Get(filename string) ([]byte, error) {
	return ioutil.ReadFile("config/" + filename)
}
