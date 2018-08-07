package main

import (
	"database/sql"
	"fmt"
)

func getDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host='%s' sslmode=disable user='%s' password='%s' dbname='%s'",
		getConf("POSTGRES_HOST"),
		getConf("POSTGRES_USER"),
		getConf("POSTGRES_PASSWORD"),
		getConf("POSTGRES_DB"),
	))
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS t1 (
			key VARCHAR(50) PRIMARY KEY,
			value integer NOT NULL DEFAULT 0
		);

		INSERT INTO t1 (key, value)
		VALUES ('counter', 0)
		ON CONFLICT DO NOTHING;`,
	); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func getCounter(db *sql.DB) (int, error) {
	row := db.QueryRow("SELECT value FROM t1 WHERE key='counter';")
	var counter int
	err := row.Scan(&counter)
	return counter, err
}

func increaseCounter(db *sql.DB) error {
	_, err := db.Exec("UPDATE t1 SET value = value + 1 WHERE key = 'counter';")
	return err
}

func resetCounter(db *sql.DB) error {
	_, err := db.Exec("UPDATE t1 SET value = 0 WHERE key = 'counter';")
	return err
}
