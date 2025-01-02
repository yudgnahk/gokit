package utils

import (
	"os"
	"strings"
)

// CreateFile ...
func CreateFile(fileName string, content []byte) error {
	_, err := os.Create(fileName)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Replace(content []byte, old string, new string) []byte {
	return []byte(strings.ReplaceAll(string(content), old, new))
}
