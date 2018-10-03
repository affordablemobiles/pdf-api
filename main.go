package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"google.golang.org/appengine"
)

func main() {
	r := mux.NewRouter()

	// Frontend pages.
	r.HandleFunc("/", indexHandler)
	//-Frontend

	// API
	api_base := mux.NewRouter()
	r.PathPrefix("/pdf/api/v1").Handler(negroni.New(
		negroni.Wrap(api_base),
	))
	api_sub := api_base.PathPrefix("/pdf/api/v1").Subrouter()

	api_sub.HandleFunc("/", indexHandler)
	// PDF Merge
	api_sub.HandleFunc("/pdf-merge", pdfMergeGETHandler).Methods("GET")
	api_sub.HandleFunc("/pdf-merge", pdfMergePOSTHandler).Methods("POST")
	//--
	// PDF Merge from GCS
	api_sub.HandleFunc("/pdf-merge-from-gcs", pdfMergeGCSGETHandler).Methods("GET")
	api_sub.HandleFunc("/pdf-merge-from-gcs", pdfMergeGCSPOSTHandler).Methods("POST")
	//--
	//-API

	// Handle all HTTP requests with our router.
	http.Handle("/", r)

	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Permission Denied", 403)
}
