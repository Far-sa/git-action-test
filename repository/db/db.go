package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type mysqlDB struct {
	db *sql.DB
}

func (m *mysqlDB) DB() *sql.DB {
	return m.db
}

const (
	retryAttempts = 3
	sleepDuration = 2 * time.Second
)

func NewMySQLDB() (*mysqlDB, error) {

	dsn := "user:password@tcp(localhost:3306)/testdb?parseTime=true"
	var db *sql.DB
	var err error

	for i := 0; i < retryAttempts; i++ { // Retry logic
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Attempt %d: failed to open database: %v", i+1, err)
			time.Sleep(sleepDuration) // Wait before retrying
			continue
		}

		err = db.Ping()
		if err == nil {
			db.SetMaxOpenConns(10)           // Example of setting max open connections
			db.SetMaxIdleConns(5)            // Example of setting max idle connections
			db.SetConnMaxLifetime(time.Hour) // Example of setting connection max lifetime
			break
		}

		log.Printf("Attempt %d: failed to ping database: %v", i+1, err)
		time.Sleep(sleepDuration) // Wait before retrying

		defer db.Close()
	}

	if err != nil {
		return nil, fmt.Errorf("could not connect to the database after multiple attempts: %w", err)
	}

	return &mysqlDB{db: db}, nil
}

func (db *mysqlDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args...)
}

func (db *mysqlDB) Get(dest interface{}, query string, args ...interface{}) error {
	return db.db.QueryRow(query, args...).Scan(dest)
}

func (db *mysqlDB) Prepare(query string) (*sql.Stmt, error) {
	return db.db.Prepare(query)
}
