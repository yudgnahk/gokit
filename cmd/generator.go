package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yudgnahk/gokit/utils"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the project based on source",
	Run: func(cmd *cobra.Command, args []string) {
		// get source path
		source, err := cmd.Flags().GetString("source")
		if err != nil {
			logrus.Fatal("Source is required")
		}
		target, err := cmd.Flags().GetString("target")
		if err != nil {
			logrus.Fatal("Target is required")
		}
		moduleName, err := cmd.Flags().GetString("module")
		if err != nil {
			logrus.Fatal("Module is required")
		}

		rootDir := Scan(source)

		// check if rootDir contains a go.mod file
		isGoProject := false
		sourceModuleName := ""
		for _, file := range rootDir.Files {
			if file.Name == "go.mod" {
				isGoProject = true
				sourceModuleName = utils.GetModuleName(fmt.Sprintf("%s/%s", rootDir.Path, file.Name))
				break
			}
		}

		if !isGoProject {
			logrus.Fatal("Source is not a Golang project")
		}

		Generate(target, rootDir, sourceModuleName, moduleName)
	},
}

func Generate(target string, dir Dir, sourceModuleName, moduleName string) {
	// check if the target folder exists and is empty
	// if not, create it
	if _, err := os.Stat(target); os.IsNotExist(err) {
		os.Mkdir(target, os.ModePerm)
	} else {
		// check if the folder is empty
		isEmpty, err := utils.IsEmpty(target)
		if err != nil {
			logrus.Fatal(err)
		}

		if !isEmpty {
			logrus.Fatal("Target folder is not empty")
		}
	}

	// create all dirs and files based on the source dir
	// loop through the dir, read file, replace module name and write new file to target
	CreateFiles(dir, target, sourceModuleName, moduleName)
	CreateDirs(dir, target, sourceModuleName, moduleName)
}

func CreateDirs(dir Dir, target, sourceModuleName, moduleName string) {
	// loop through the dir, create dirs
	for _, d := range dir.Dirs {
		// create dir
		err := os.Mkdir(fmt.Sprintf("%s/%s", target, d.Name), os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}

		// create files
		CreateFiles(d, fmt.Sprintf("%s/%s", target, d.Name), sourceModuleName, moduleName)

		// call CreateDirs recursively
		CreateDirs(d, fmt.Sprintf("%s/%s", target, d.Name), sourceModuleName, moduleName)
	}
}

func CreateFiles(dir Dir, target, sourceModuleName, moduleName string) {
	// loop through the dir, read file, replace module name and write new file to target
	for _, file := range dir.Files {
		// read file
		content, err := os.ReadFile(fmt.Sprintf("%s/%s", dir.Path, file.Name))
		if err != nil {
			logrus.Fatal(err)
		}

		// replace module name
		content = utils.Replace(content, sourceModuleName, moduleName)
		// replace the last part of the source module name with the last part of the module name
		content = utils.Replace(content, utils.GetLastPart(sourceModuleName), utils.GetLastPart(moduleName))

		// write file
		err = os.WriteFile(fmt.Sprintf("%s/%s", target, file.Name), content, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func init() {
	RootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("source", "s", "", "Source path")
	generateCmd.Flags().StringP("target", "t", "", "Target path")
	generateCmd.Flags().StringP("module", "m", "", "Module name")
}
