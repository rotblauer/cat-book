package main

import (
	"net/http"
	"strconv"
	"flag"
)

func main() {
	var porty int
	flag.IntVar(&porty, "port", 8080, "port to serve and protect")

	if bolterr := InitBoltDB(); bolterr == nil {
		defer GetDB().Close()
	}

	router := CatRouter()

	http.Handle("/", router)

	http.ListenAndServe(":"+strconv.Itoa(porty), nil)
}
