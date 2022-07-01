package main

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yudgnahk/gokit/cmd"
)

func main() {
	viper.AutomaticEnv()
	pwd, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return
	}

	_, err = filepath.EvalSymlinks(pwd)
	if err != nil {
		logrus.Error(err)
		return
	}

	cmd.Execute()
}
