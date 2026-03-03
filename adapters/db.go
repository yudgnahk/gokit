package database

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// Errs alias
var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type DBType int

const (
	MySQLType DBType = iota
	PostgreSQLType
)

// DBAdapter interface represent adapter connect to DB
type DBAdapter interface {
	Open(connectionString string, dbType DBType, schemaName string) error
	Close()
	Begin() DBAdapter
	RollbackUselessCommitted()
	Commit()
	Gormer() *gorm.DB
	DB() *sql.DB
	Schema() string
}

type adapter struct {
	gormer      *gorm.DB
	isCommitted bool
	dbType      DBType
	dbName      string
}

// NewDB returns a new instance of DB.
func NewDB() DBAdapter {
	return &adapter{}
}

// Open opens a DB connection.
func (db *adapter) Open(connectionString string, dbType DBType, schemaName string) error {
	db.dbType = dbType
	db.dbName = schemaName

	var dialect gorm.Dialector
	switch dbType {
	case MySQLType:
		dialect = mysql.Open(connectionString)
	case PostgreSQLType:
		dialect = postgres.Open(connectionString)
	}

	gormDB, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return err
	}
	db.gormer = gormDB

	gormDB.Use(
		dbresolver.Register(dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)
	return nil
}

// Close closes DB connection.
func (db *adapter) Close() {
	_ = db.DB().Close()
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
	database, _ := db.gormer.DB()
	return database
}

func (db *adapter) Schema() string {
	return db.dbName
}
