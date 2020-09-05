package utils

import (
	"os/exec"
)

func GitInit(dir string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", "git init")
	cmd.Dir = dir
	return cmd.CombinedOutput()
}
