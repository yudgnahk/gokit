package repositories

import (
	"fmt"

	database "github.com/yudgnahk/gokit/adapters"
	"github.com/yudgnahk/gokit/configs"
	"github.com/yudgnahk/gokit/models"
)

type Schema interface {
	GetTables() ([]string, error)
	GetColumns(tableName string) ([]*models.MySQLColumn, error)
}

type schema struct {
	db database.DBAdapter
}

const (
	BaseTable   = "BASE TABLE"
	TableType   = "table_type"
	TableSchema = "table_schema"
	TableName   = "table_name"
	Tables      = "information_schema.tables"
	Columns     = "information_schema.columns"
)

func NewSchema(db database.DBAdapter) Schema {
	return &schema{db: db}
}

func (r *schema) GetTables() ([]string, error) {
	gormer := r.db.Gormer()
	var tables []string

	if err := gormer.Table(Tables).
		Where(fmt.Sprintf("%v = ? AND %v = ?", TableType, TableSchema), BaseTable, configs.AppConfig.DB.Database).
		Pluck(TableName, &tables).Error; err != nil {
		return tables, err
	}

	return tables, nil
}

func (r *schema) GetColumns(tableName string) ([]*models.MySQLColumn, error) {
	gormer := r.db.Gormer()
	var columns []*models.MySQLColumn
	fmt.Println(configs.AppConfig.DB.Database, tableName)
	if err := gormer.Model(&models.MySQLColumn{}).
		Select("column_name, ordinal_position, data_type, column_type, column_key, extra").
		Where(fmt.Sprintf("%v = ? AND %v = ?", TableSchema, TableName),
			configs.AppConfig.DB.Database, tableName).
		Find(&columns).Error; err != nil {
		return columns, err
	}
	return columns, nil
}
