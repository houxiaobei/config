package config

import "fmt"

type FileExtError struct {
	Ext string
	Err error
}

func (f *FileExtError) Error() string {
	return fmt.Sprintf("ext %s is err, json or yaml ", f.Ext)
}

type FileNotFound struct {
	File string
	Err error
}

func (f *FileNotFound) Error() string {
	return fmt.Sprintf("file %s not found", f.File)
}