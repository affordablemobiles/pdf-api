package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"cloud.google.com/go/errorreporting"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	exporter, err := stackdriver.NewExporter(stackdriver.Options{ProjectID: gae_project()})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)
	defer exporter.Flush()

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

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

	//---
	// Error reporting...
	//---
	errorClient, err := errorreporting.NewClient(context.Background(), gae_project(), errorreporting.Config{
		ServiceName:    gae_service(),
		ServiceVersion: gae_version(),
		OnError: func(err error) {
			log.Printf("Could not log error: %v", err)
		},
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer errorClient.Close()

	recover := negroni.NewRecovery()
	recover.PrintStack = true
	recover.LogStack = false
	recover.PanicHandlerFunc = func(info *negroni.PanicInformation) {
		errorClient.Report(errorreporting.Entry{
			Req:   info.Request,
			Error: info.RecoveredPanic.(error),
			Stack: info.Stack,
		})
	}
	recover.Formatter = &CustomPanicFormatter{}

	// Handle all HTTP requests with our router.
	http.Handle("/", negroni.New(
		recover,
		negroni.Wrap(&ochttp.Handler{
			Propagation: &propagation.HTTPFormat{},
			Handler: r,
}		),
	))

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

type CustomPanicFormatter struct{}

func (t *CustomPanicFormatter) FormatPanicError(w http.ResponseWriter, r *http.Request, infos *negroni.PanicInformation) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func gae_project() string {
	return os.Getenv("GOOGLE_CLOUD_PROJECT")
}

func gae_service() string {
	if svc := os.Getenv("GAE_SERVICE"); svc != "" {
		return svc
	} else {
		return os.Getenv("K_SERVICE")
	}
}

func gae_version() string {
	if ver := os.Getenv("GAE_VERSION"); ver != "" {
		return ver
	} else {
		return os.Getenv("K_REVISION")
	}
}