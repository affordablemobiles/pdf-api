package main

import (
	"net/http"

	glog "github.com/a1comms/go-gaelog/v2"
)

func error500Handler(w http.ResponseWriter, r *http.Request, err error) {
	ctx := glog.GetContext(r)

	glog.Errorf(ctx, nil, "%s", err)

	http.Error(w, "Internal Server Error", 500)
}
