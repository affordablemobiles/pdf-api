package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	// Frontend pages.
	r.NotFoundHandler = r.NewRoute().BuildOnly().HandlerFunc(defaultHandler).GetHandler()
	//-Frontend

	// API
	api_base := mux.NewRouter()
	r.PathPrefix("/pdf/api/v1").Handler(negroni.New(
		negroni.Wrap(api_base),
	))
	api_sub := api_base.PathPrefix("/pdf/api/v1").Subrouter()

	api_sub.NotFoundHandler = r.NewRoute().BuildOnly().HandlerFunc(defaultHandler).GetHandler()

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

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Permission Denied", 403)
}
