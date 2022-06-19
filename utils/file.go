package utils

import "os"

func NewFile() *file {
	return &file{}
}

type file struct {
}

func (f *file) PathIsExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if ok := os.IsNotExist(err); ok {
		return false, err
	}
	return false, err
}
