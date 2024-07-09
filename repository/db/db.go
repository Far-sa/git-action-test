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

func NewMySQLDB(dataSourceName string) (*mysqlDB, error) {
	var db *sql.DB
	var err error

	for i := 0; i < 3; i++ { // Retry logic
		db, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Printf("Attempt %d: failed to open database: %v", i+1, err)
			time.Sleep(2 * time.Second) // Wait before retrying
			continue
		}

		err = db.Ping()
		if err == nil {
			break
		}

		log.Printf("Attempt %d: failed to ping database: %v", i+1, err)
		db.Close()                  // Ensure the database connection is closed if ping fails
		time.Sleep(2 * time.Second) // Wait before retrying
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
