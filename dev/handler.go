package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func rootHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "Hello World")
		writeFooter(w)
	}
}

func counterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		counter, err := getCounter(db)
		if err != nil {
			writeError(w, err)
			return
		}
		if err := increaseCounter(db); err != nil {
			writeError(w, err)
			return
		}
		fmt.Fprintln(w, "Counter is "+strconv.Itoa(counter))
		writeFooter(w)
	}
}
func resetHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if err := resetCounter(db); err != nil {
			writeError(w, err)
			return
		}
		fmt.Fprintln(w, "Counter is reset to 0")
		writeFooter(w)
	}
}

func writeFooter(w io.Writer) {
	hostname, _ := os.Hostname()
	fmt.Fprintln(w)
	fmt.Fprintln(w, "-----------")
	fmt.Fprintln(w, "Processed by: "+hostname)
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	fmt.Println("ERROR: " + err.Error())
	writeFooter(w)
}
