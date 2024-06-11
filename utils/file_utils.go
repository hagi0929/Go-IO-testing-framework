package utils

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteFile(path string, data string) error {
	err := ioutil.WriteFile(path, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}
