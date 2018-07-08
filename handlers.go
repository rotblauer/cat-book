package main

import (
	"encoding/json"
	"image"
	"net/http"
	"github.com/gorilla/mux"
)
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


var routes = Routes{Route{
	"CatStore",
	"Post",
	"/",
	catPic,
},
}

func catPic(w http.ResponseWriter, r *http.Request) {
	var catPic image.Image

	if r.Body == nil {
		http.Error(w, "Please send a request body", 500)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&catPic)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func CatRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
