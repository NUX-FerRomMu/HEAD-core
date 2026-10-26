// Package store opens local SQLite databases. Feature packages own their tables.
package store

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, fmt.Errorf("create state directory: %w", err)
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("open state: %w", err)
	}
	if err = f.Close(); err != nil {
		return nil, err
	}
	// Encode path punctuation so it cannot become SQLite URI parameters.
	absolute, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	uri := &url.URL{Scheme: "file", Path: absolute}
	db, err := sql.Open("sqlite", uri.String())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	// Keep connection-local settings for the lifetime of this short-lived command.
	db.SetMaxIdleConns(1)
	for _, statement := range []string{"PRAGMA busy_timeout=5000", "PRAGMA foreign_keys=ON", "PRAGMA journal_mode=WAL"} {
		if _, err = db.Exec(statement); err != nil {
			db.Close()
			return nil, fmt.Errorf("initialize state: %w", err)
		}
	}
	return db, nil
}
