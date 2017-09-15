package main

import (
	"time"
	"testing"
	"database/sql"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

func catch(t *testing.T) {
	err := recover()
	if err != nil {
		t.Error(err)
	}
}

func TestUnsetLifetime(t *testing.T) {
	dsn := "root:test@tcp(127.0.0.1:33306)/test?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dsn)
	print(t, err)
	
	defer db.Close()
	
	print(t, db.Ping())
	
	// Wait for 5 seconds. This should be enough to timeout the conn, since `wait_timeout` is 3s
	time.Sleep(5 * time.Second)
	defer catch(t)
	
	// Simply attempt to begin a transaction
	tx, err := db.Begin()
	print(t, err)
	defer tx.Rollback()
}

func TestSetLifetime(t *testing.T) {
	dsn := "root:test@tcp(127.0.0.1:33306)/test?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dsn)
	print(t, err)
	db.SetConnMaxLifetime(2)
	
	defer db.Close()
	
	print(t, db.Ping())
	
	// Wait for 5 seconds. This should be enough to timeout the conn, since `wait_timeout` is 3s
	time.Sleep(5 * time.Second)
	defer catch(t)
	
	// Simply attempt to begin a transaction
	tx, err := db.Begin()
	print(t, err)
	defer tx.Rollback()
}

func print(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
