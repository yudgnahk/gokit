package constants

import (
	"gorm.io/datatypes"
)

// Definition constants
const (
	ModuleName          = "MODULE_NAME"
	BackStick           = "BACK_STICK"
	BasePath            = "BASE_PATH"
	ControllerName      = "CONTROLLER_NAME"
	ControllerNameLower = "controller_name"
	ServiceName         = "SERVICE_NAME"
	ServiceNameLower    = "service_name"
	RepositoryName      = "REPOSITORY_NAME"
	RepositoryNameLower = "repository_name"
	ModelNameLower      = "model_name"
	ModelContent        = "MODEL_CONTENT"
	ColumnName          = "COLUMN_NAME"
	ColumnNameLower     = "column_name"
	ColumnType          = "column_type"

	Space       = " "
	Underscore  = "_"
	EmptyString = ""

	// MySQLDataTypes
	DateTime  = "datetime"
	Varchar   = "varchar"
	Bigint    = "bigint"
	Longtext  = "longtext"
	Tinyint   = "tinyint"
	Int       = "int"
	Enum      = "enum"
	Smallint  = "smallint"
	Double    = "double"
	Json      = "json"
	Timestamp = "timestamp"
	Text      = "text"

	// GolangTypes
	Int64   = "int64"
	Bool    = "bool"
	String  = "string"
	Time    = "time.Time"
	Float64 = "float64"
	JSON    = "datatypes.JSON"
)

type Test struct {
	JSON datatypes.JSON
}
