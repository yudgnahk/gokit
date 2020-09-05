package templates

// DBAdapterTemplate ...
const DBAdapterTemplate = `package database

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"

	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Errs alias
var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

// Package constants definition.
const (
	MySQLDialect = "mysql"
)

// DBAdapter interface represent adapter connect to DB
type DBAdapter interface {
	Open(string) error
	Close()
	Begin() DBAdapter
	RollbackUselessCommitted()
	Commit()
	Gormer() *gorm.DB
	DB() *sql.DB
}

type adapter struct {
	gormer      *gorm.DB
	isCommitted bool
}

// NewDB returns a new instance of DB.
func NewDB() DBAdapter {
	return &adapter{}
}

// Open opens a DB connection.
func (db *adapter) Open(connectionString string) error {
	gormDB, err := gorm.Open(MySQLDialect, connectionString)
	if err != nil {
		return err
	}
	gormDB.DB().SetConnMaxLifetime(time.Minute * 10)
	db.gormer = gormDB
	return nil
}

// Close closes DB connection.
func (db *adapter) Close() {
	_ = db.gormer.Close()
}

// Begin starts a DB transaction.
func (db *adapter) Begin() DBAdapter {
	tx := db.gormer.Begin()
	return &adapter{
		gormer:      tx,
		isCommitted: false,
	}
}

// RollbackUselessCommitted rollbacks useless DB transaction committed.
func (db *adapter) RollbackUselessCommitted() {
	if !db.isCommitted {
		db.gormer.Rollback()
	}
}

// Commit commits a DB transaction.
func (db *adapter) Commit() {
	if !db.isCommitted {
		db.gormer.Commit()
		db.isCommitted = true
	}
}

// Gormer returns an instance of gorm.DB.
func (db *adapter) Gormer() *gorm.DB {
	return db.gormer
}

// DB returns an instance of sql.DB.
func (db *adapter) DB() *sql.DB {
	return db.gormer.DB()
}
`
