package main

import (
	"flag"
	"net/http"
	"strconv"
)

func main() {
	var porty int
	flag.IntVar(&porty, "port", 3000, "port to serve and protect")

	if bolterr := InitBoltDB(); bolterr == nil {
		defer GetDB().Close()
	}

	router := CatRouter()

	http.Handle("/", router)

	http.ListenAndServe(":"+strconv.Itoa(porty), nil)
}
