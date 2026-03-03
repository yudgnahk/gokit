package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/yudgnahk/gokit/constants"
)

// GoModInit ...
func GoModInit(moduleName, dir string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("go mod init %v", moduleName))
	cmd.Dir = dir
	return cmd.CombinedOutput()
}

// GoModTidy ...
func GoModTidy(dir string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", "go mod tidy")
	cmd.Dir = dir
	return cmd.CombinedOutput()
}

func GoFmt(dir string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", "gofmt -w .")
	cmd.Dir = dir
	return cmd.CombinedOutput()
}

// GetModuleName ...
func GetModuleName(path string) string {
	if path == "" {
		path = "go.mod"
	}
	var result string
	var count = 0
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
		firstRow := scanner.Text()
		parts := strings.Split(firstRow, constants.Space)
		result = parts[len(parts)-1]

		if count > 0 {
			return result
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
