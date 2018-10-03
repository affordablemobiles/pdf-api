package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func error500Handler(w http.ResponseWriter, r *http.Request, err error) {
	ctx := appengine.NewContext(r)

	log.Errorf(ctx, "%s", err)

	http.Error(w, "Internal Server Error", 500)
}
