package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yudgnahk/gokit/constants"
	"github.com/yudgnahk/gokit/templates"
	"github.com/yudgnahk/gokit/utils"
)

var projectName, moduleName, basePath string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"n"},
	Short:   "Init project",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(constants.ColorGreen, "Project name: ")
		if scanner.Scan() {
			projectName = scanner.Text()
		}
		fmt.Print("Module name: ")
		if scanner.Scan() {
			moduleName = scanner.Text()
		}

		fmt.Print("Base path: ")
		if scanner.Scan() {
			basePath = scanner.Text()
		}

		fmt.Println(constants.ColorYellow, "======Initialize=====")

		rootFolder := projectName
		folders := []string{"cmd", "adapters", "migrations", "configs", "errors", "dtos", "utils", "controllers", "repositories", "services"}
		fmt.Println("Create base folders...")
		time.Sleep(time.Second)
		err := os.Mkdir(rootFolder, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}

		err = os.Chdir(rootFolder)
		if err != nil {
			logrus.Fatal(err)
		}

		for i := range folders {
			err = os.Mkdir(folders[i], os.ModePerm)
			if err != nil {
				logrus.Fatal(err)
			}
		}

		params := utils.StandardizeParams{
			ModuleName: moduleName,
			BasePath:   basePath,
		}

		fmt.Println("Create base files...")
		time.Sleep(time.Second)
		err = utils.CreateFile(fmt.Sprintf("./%v/%v", "cmd", "main.go"), []byte(utils.StandardizedTemplate(templates.MainContent, params)))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v/%v", "controllers", "base.go"), []byte(utils.StandardizedTemplate(templates.BaseController, params)))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v/%v", "controllers", "app.go"), []byte(utils.StandardizedTemplate(templates.AppController, params)))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v/%v", "adapters", "db.go"), []byte(utils.StandardizedTemplate(templates.DBAdapterTemplate, params)))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v/%v", "dtos", "base.go"), []byte(utils.StandardizedTemplate(templates.BaseDtos, params)))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v/%v", "dtos", "health.go"), []byte(utils.StandardizedTemplate(templates.HealthDtos, params)))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v/%v", "errors", "errors.go"), []byte(utils.StandardizedTemplate(templates.ErrorsTemplate, params)))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v/%v", "errors", "codes.go"), []byte(utils.StandardizedTemplate(templates.ErrorCodesTemplate, params)))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v/%v", "utils", "response.go"), []byte(utils.StandardizedTemplate(templates.ResponseUtil, params)))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v/%v", "configs", "config.go"), []byte(utils.StandardizedTemplate(templates.ConfigTemplate, params)))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v", "Makefile"), []byte(templates.MakefileTemplate))
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.CreateFile(fmt.Sprintf("./%v", ".env"), []byte(templates.EnvTemplate))
		if err != nil {
			logrus.Fatal(err)
		}

		fmt.Println("Go mod init...")
		time.Sleep(time.Second)

		folder, _ := os.Getwd()
		output, err := utils.GoModInit(moduleName, folder)
		if err != nil {
			logrus.Fatalf("Error: %v, Output: %v", err, output)
		}

		output, err = utils.GoModTidy(folder)
		if err != nil {
			logrus.Fatalf("Error: %v, Output: %v", err, output)
		}

		fmt.Println("Git init...")
		time.Sleep(time.Second)

		output, err = utils.GitInit(folder)
		if err != nil {
			logrus.Fatalf("Error: %v, Output: %v", err, output)
		}

		fmt.Println("Finish initialization!")
		time.Sleep(time.Second / 4)
		fmt.Print(constants.ColorBlue, fmt.Sprintf("Move to your work by: $cd %v", projectName))
		fmt.Println()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
