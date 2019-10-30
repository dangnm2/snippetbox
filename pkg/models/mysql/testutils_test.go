package mysql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type config struct {
	user     string
	password string
	name     string
	address  string
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getDBConfig() config {
	return config{
		name:     getEnv("SNIPPETBOX_DB_NAME", "snippetbox"),
		user:     getEnv("SNIPPETBOX_DB_USER", "root"),
		password: getEnv("SNIPPETBOX_DB_PASSWORD", ""),
		address:  getEnv("SNIPPETBOX_DB_ADDRESS", "127.0.0.1"),
	}
}

func newTestDB(t *testing.T) (*sql.DB, func()) {
	config := getDBConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true", config.user, config.password, config.address, config.name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}

	script, err := ioutil.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))

	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		script, err := ioutil.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()

	}

}
