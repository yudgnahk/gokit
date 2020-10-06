package templates

type Money struct {
	ID int64 `gorm:"column(id)"`
}

const ModelTemplate = `package models

type model_name struct {
	MODEL_CONTENT
}
`

const GormTemplate = `COLUMN_NAME column_type BACK_STICKgorm:"column(column_name)"BACK_STICK`

const ColumnTemplate = `COLUMN_NAME column_type`
