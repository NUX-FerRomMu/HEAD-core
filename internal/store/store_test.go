package store

import (
	"path/filepath"
	"testing"
)

func TestReopen(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested?literal#dir", "state.sqlite")
	db, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = db.Exec("CREATE TABLE sample (id INTEGER); INSERT INTO sample VALUES (42)"); err != nil {
		t.Fatal(err)
	}
	db.Close()
	db, err = Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	var value int
	if err = db.QueryRow("SELECT id FROM sample").Scan(&value); err != nil || value != 42 {
		t.Fatalf("%d %v", value, err)
	}
	var fk int
	if err = db.QueryRow("PRAGMA foreign_keys").Scan(&fk); err != nil || fk != 1 {
		t.Fatalf("foreign keys: %d %v", fk, err)
	}
}
