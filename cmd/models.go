package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	database "github.com/yudgnahk/gokit/adapters"
	"github.com/yudgnahk/gokit/configs"
	"github.com/yudgnahk/gokit/constants"
	"github.com/yudgnahk/gokit/models"
	"github.com/yudgnahk/gokit/repositories"
	"github.com/yudgnahk/gokit/templates"
	"github.com/yudgnahk/gokit/utils"
)

// genModelCmd represents the new command
var genModelsCmd = &cobra.Command{
	Use:     "models",
	Aliases: []string{"m"},
	Short:   "Generate models from database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(constants.ColorBlue, "Generate models...\n")

		dir := "models"

		os.Mkdir(dir, os.ModePerm)

		db := database.NewDB()
		if err := db.Open(configs.AppConfig.DB.ConnectionString()); err != nil {
			logrus.Fatalf("Creating connection to DB: %v", err)
		}

		schema := repositories.NewSchema(db)
		tables, err := schema.GetTables()
		if err != nil {
			logrus.Fatalf("error: %v", err)
		}

		for i := range tables {
			if tables[i] == "migrations" {
				continue
			}
			columns, err := schema.GetColumns(tables[i])
			if err != nil {
				logrus.Errorf("Get columns of table %v got error %v", tables[i], err)
			} else {
				fmt.Println("Creating model", utils.Camel(tables[i], false), "...")
				createModel(columns, dir, tables[i])
			}
		}

		utils.GoFmt(dir)
		fmt.Println("Finish generating models")
	},
}

func createModel(columns []*models.MySQLColumn, dir, tableName string) {
	fileName := fmt.Sprintf("./models/%v.go", utils.Snake(tableName))
	columnsRaw := ""

	for i := 0; i < len(columns); i++ {
		column := templates.ColumnTemplate
		columnNameCamel := utils.Camel(columns[i].ColumnName, false)
		dataType := getDataType(columns[i].ColumnType)

		column = standardizedColumn(column, columnNameCamel, dataType)
		columnsRaw += fmt.Sprintf("\n%v", column)
	}

	model := strings.ReplaceAll(templates.ModelTemplate, constants.ModelContent, columnsRaw)
	model = strings.ReplaceAll(model, constants.ModelNameLower, utils.Camel(tableName, false))
	err := utils.CreateFile(fileName, []byte(model))
	if err != nil {
		logrus.Fatal(err)
	}
}

func standardizedColumn(column, columnNameCamel, dataType string) string {
	column = strings.ReplaceAll(column, constants.ColumnName, columnNameCamel)
	column = strings.ReplaceAll(column, constants.ColumnType, dataType)
	return column
}

func getDataType(sqlType string) string {
	if strings.Contains(sqlType, constants.Tinyint) {
		return constants.Bool
	}

	if strings.Contains(sqlType, constants.Int) ||
		strings.Contains(sqlType, constants.Bigint) ||
		strings.Contains(sqlType, constants.Smallint) {
		return constants.Int64
	}

	if strings.Contains(sqlType, constants.Varchar) ||
		strings.Contains(sqlType, constants.Text) ||
		strings.Contains(sqlType, constants.Longtext) {
		return constants.String
	}

	if strings.Contains(sqlType, constants.Double) {
		return constants.Float64
	}

	if strings.Contains(sqlType, constants.DateTime) ||
		strings.Contains(sqlType, constants.Timestamp) {
		return constants.Time
	}

	if strings.Contains(sqlType, constants.Json) {
		return constants.JSON
	}

	if strings.Contains(sqlType, constants.Enum) {
		return "enum"
	}

	return "undefined"
}

func init() {
	_, err := configs.New()
	if err != nil {
		os.Exit(99)
	}

	RootCmd.AddCommand(genModelsCmd)
}
