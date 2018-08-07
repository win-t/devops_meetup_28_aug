package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/payfazz/go-middleware"
	"github.com/payfazz/go-middleware/common"
	"github.com/payfazz/go-router/path"
)

func main() {
	db, err := getDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	handler := middleware.Compile(
		common.BasicPack(),
		path.C(path.H{
			"/":        rootHandler(db),
			"/counter": counterHandler(db),
			"/reset":   resetHandler(db),
		}),
	)

	panic(http.ListenAndServe(getConf("APP_ADDR"), handler))
}
