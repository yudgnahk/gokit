package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/khanghldk/gokit/constants"
	"github.com/khanghldk/gokit/templates"
	"github.com/khanghldk/gokit/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n"},
	Short:   "Some useful generators",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			logrus.Fatal("There are not enough params")
		}

		var genType = args[0]

		var name = strings.Join(args[1:], constants.Underscore)
		moduleName := utils.GetModuleName()

		params := utils.StandardizeParams{
			ModuleName:     moduleName,
			ControllerName: name,
			ServiceName:    name,
			RepositoryName: name,
		}

		switch genType {
		case "controller":
			err := GenerateController(params)
			if err != nil {
				logrus.Fatal(err)
			}

			err = GenerateService(params)
			if err != nil {
				logrus.Fatal(err)
			}

			err = GenerateRepository(params)
			if err != nil {
				logrus.Fatal(err)
			}
		case "service":
			err := GenerateService(params)
			if err != nil {
				logrus.Fatal(err)
			}

			err = GenerateRepository(params)
			if err != nil {
				logrus.Fatal(err)
			}
		case "repository":
			err := GenerateRepository(params)
			if err != nil {
				logrus.Fatal(err)
			}
		}
	},
}

func GenerateController(params utils.StandardizeParams) error {
	path := fmt.Sprintf("./%v", "controllers")
	fileName := utils.Snake(params.ControllerName)

	os.Mkdir(path, os.ModePerm)

	content := utils.StandardizedTemplate(templates.ControllerTemplate, params)
	filePath := fmt.Sprintf("%v/%v.go", path, fileName)

	err := utils.CreateFile(filePath, []byte(content))
	return err
}

func GenerateService(params utils.StandardizeParams) error {
	path := fmt.Sprintf("./%v", "services")
	fileName := utils.Snake(params.ServiceName)
	os.Mkdir(path, os.ModePerm)

	filePath := fmt.Sprintf("%v/%v.go", path, fileName)

	content := utils.StandardizedTemplate(templates.ServiceTemplate, params)
	err := utils.CreateFile(filePath, []byte(content))
	return err
}

func GenerateRepository(params utils.StandardizeParams) error {
	path := fmt.Sprintf("./%v", "repositories")
	fileName := utils.Snake(params.RepositoryName)
	os.Mkdir(path, os.ModePerm)

	filePath := fmt.Sprintf("%v/%v.go", path, fileName)

	content := utils.StandardizedTemplate(templates.RepositoryTemplate, params)
	err := utils.CreateFile(filePath, []byte(content))
	return err
}

func init() {
	RootCmd.AddCommand(newCmd)
}
