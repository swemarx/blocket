package main

import (
	//"os"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func openSqlite(filename string) (*sql.DB, error) {
	//dsn := fmt.Sprintf("file:%s?cache=shared", filename)
	//db, err := sql.Open("sqlite3", dsn)
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = checkSqlite(db)
	if err != nil {
		fmt.Println(err)
		closeSqlite(db)
		return nil, err
	}

	return db, nil
}

func checkSqlite(db *sql.DB) error {
	tableQueryTmpl := "SELECT name FROM sqlite_master WHERE type='table' AND name='%s'"
	tableQuery := fmt.Sprintf(tableQueryTmpl, "users")
	rows, err := db.Query(tableQuery)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()
	var tableName string
	err = rows.Scan(&tableName)
	if err != nil && config.Debug {
		fmt.Println("[debug] Found table \"users\"")
		return nil
	} else if err == nil && config.Debug {
		return errors.New("[debug] Did not find table \"users\"")
	}

	return nil
}

func closeSqlite(db *sql.DB) {
	db.Close()
}
