package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DatabaseAdapter struct {
	DB *sql.DB
}

var dbConn = &DatabaseAdapter{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

func NewAdapter(dataSourceName string) (*DatabaseAdapter, error) {
	// connect
	d, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}
	dbConn.DB = d
	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	fmt.Println("*** Pinged database successfully! ***")
	return nil
}

func (da DatabaseAdapter) CloseDbConnection() {
	err := da.DB.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}
