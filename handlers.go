package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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
	"POST",
	"/CatStore",
	catPic,
}, Route{
	"CatPic",
	"GET",
	"/{db}/{z}/{x}/{y}",
	catPic,
},
}

/// smorganizing

// get cat image + trackerpt
// store image in db, or on hd with https://github.com/deiwin/imstor and db manages trackpoint -> file mapper,

/// imstor nice since can do small/large/original auto pics

func getCatPics(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// store a cat pic to db
// TODO geograph awares with trackpoint mapping
func catPic(w http.ResponseWriter, r *http.Request) {
	var cat CatPic
	println(r.Body)
	if r.Body == nil {
		http.Error(w, "Please send a request body", 500)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	storeCat(cat)

}

func CatRouter() *mux.Router {
	println("got router")
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		println("got route")

	}

	return router
}
