package utils

import (
	"io/ioutil"
	"os"
)

func CreateFile(fileName string, content []byte) error {
	_, err := os.Create(fileName)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		return err
	}

	return nil
}
