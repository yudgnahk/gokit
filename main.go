package main

import (
	"os"
	"path/filepath"

	"gokit/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	pwd, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return
	}

	pwd, err = filepath.EvalSymlinks(pwd)
	if err != nil {
		logrus.Error(err)
		return
	}

	cmd.Execute()
}
