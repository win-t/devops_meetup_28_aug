package main

import "os"

var defaultConf = map[string]string{
	"APP_ADDR":          ":8080",
	"POSTGRES_HOST":     "localhost",
	"POSTGRES_USER":     "testuser",
	"POSTGRES_PASSWORD": "testpassword",
	"POSTGRES_DB":       "testdb",
}

func getConf(key string) string {
	retval := os.Getenv(key)
	if retval == "" {
		retval = defaultConf[key]
	}
	return retval
}
