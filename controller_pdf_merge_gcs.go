package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	glog "github.com/a1comms/go-gaelog/v2"
	"github.com/mattetti/filebuffer"

	"go.opencensus.io/trace"
)

func pdfMergeGCSGETHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("<form action=\"/pdf/api/v1/pdf-merge-from-gcs\" method=\"POST\" enctype=\"application/x-www-form-urlencoded\"><textarea name=\"fdf\"></textarea><input type=\"text\" name=\"pdf_filename\" /><input type=\"submit\" value=\"Submit\" />\r\n"))
}

func pdfMergeGCSPOSTHandler(w http.ResponseWriter, r *http.Request) {
	ctx := glog.GetContext(r)

	pdfFilename := r.FormValue("pdf_filename")
	if pdfFilename == "" {
		error500Handler(w, r, fmt.Errorf("No PDF Filename Found in POST"))
		return
	}
	fdfEncData := r.FormValue("fdf")
	if fdfEncData == "" {
		error500Handler(w, r, fmt.Errorf("No FDF Data Found in POST"))
		return
	}
	base64Output := r.FormValue("base64")

	_, span := trace.StartSpan(ctx, "base64/decode")
	fdfData, err := base64.StdEncoding.DecodeString(fdfEncData)
	if err != nil {
		error500Handler(w, r, fmt.Errorf("Error decoding FDF: %s", err))
		return
	}
	span.End()

	_, span = trace.StartSpan(ctx, "fdf/strip")
	fdfData = stripFDF(fdfData)
	span.End()

	_, span = trace.StartSpan(ctx, "storage/init")
	client, err := storage.NewClient(ctx)
	if err != nil {
		error500Handler(w, r, fmt.Errorf("Error getting GCS Client: %s", err))
		return
	}
	defer client.Close()
	span.End()

	bucketName := os.Getenv("GCLOUD_STORAGE_BUCKET")
	if bucketName == "" {
		error500Handler(w, r, fmt.Errorf("Failed to get default GCS bucket name: %s", err))
		return
	}

	bucket := client.Bucket(bucketName)

	_, span = trace.StartSpan(ctx, "storage/open")
	obj, err := bucket.Object("pdf/" + pdfFilename).NewReader(ctx)
	if err != nil {
		error500Handler(w, r, fmt.Errorf("Error error opening reader for file (%s): %s", pdfFilename, err))
		return
	}
	span.End()

	_, span = trace.StartSpan(ctx, "storage/ReadAll")
	pdfData, err := ioutil.ReadAll(obj)
	if err != nil {
		error500Handler(w, r, fmt.Errorf("Error reading file (%s): %s", pdfFilename, err))
		return
	}
	span.End()

	_, span = trace.StartSpan(ctx, "storage/close")
	if err = obj.Close(); err != nil {
		error500Handler(w, r, fmt.Errorf("Error closing reader for file (%s): %s", pdfFilename, err))
		return
	}
	span.End()

	_, span = trace.StartSpan(ctx, "fdf/merge")
	output := filebuffer.New([]byte{})
	err = fdfMerge(bytes.NewReader(pdfData), bytes.NewReader(fdfData), output)
	if err != nil {
		error500Handler(w, r, fmt.Errorf("PDF/FDF Merge Failed: %s", err))
		return
	}
	span.End()
	output.Seek(0, 0)

	_, span = trace.StartSpan(ctx, "encode/write")
	defer span.End()
	if base64Output != "" {
		w.Header().Add("Content-type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		w.Write([]byte(base64.StdEncoding.EncodeToString(output.Bytes())))
	} else {
		w.Header().Add("Content-type", "application/pdf; charset=utf-8")
		w.WriteHeader(200)
		w.Write(output.Bytes())
	}
}
