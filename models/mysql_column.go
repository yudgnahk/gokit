package models

const (
	Columns = "information_schema.columns"
)

type MySQLColumn struct {
	ColumnName             string
	OrdinalPosition        int64
	ColumnDefault          int64
	IsNullable             bool
	DataType               string
	CharacterMaximumLength int64
	CharacterOctetLength   int64
	NumericPrecision       int64
	NumericScale           int64
	DatetimePrecision      int64
	CharacterSetName       string
	CollationName          string
	ColumnType             string
	ColumnKey              string
	Extra                  string
	Privileges             string
	ColumnComment          string
}

func (MySQLColumn) TableName() string {
	return Columns
}
