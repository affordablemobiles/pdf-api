package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/mattetti/filebuffer"
)

func pdfMergeGETHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte("<form action=\"/pdf/api/v1/pdf-merge\" method=\"POST\" enctype=\"application/x-www-form-urlencoded\"><textarea name=\"fdf\"></textarea><textarea name=\"pdf\"></textarea><input type=\"submit\" value=\"Submit\" />\r\n"))
}

func pdfMergePOSTHandler(w http.ResponseWriter, r *http.Request) {
	pdfEncData := r.FormValue("pdf")
	if pdfEncData == "" {
		error500Handler(w, r, fmt.Errorf("No PDF Data Found in POST"))
		return
	}
	fdfEncData := r.FormValue("fdf")
	if fdfEncData == "" {
		error500Handler(w, r, fmt.Errorf("No FDF Data Found in POST"))
		return
	}

	pdfData, err := base64.StdEncoding.DecodeString(pdfEncData)
	if err != nil {
		error500Handler(w, r, fmt.Errorf("Error decoding PDF: %s", err))
		return
	}
	fdfData, err := base64.StdEncoding.DecodeString(fdfEncData)
	if err != nil {
		error500Handler(w, r, fmt.Errorf("Error decoding FDF: %s", err))
		return
	}
	fdfData = stripFDF(fdfData)

	output := filebuffer.New([]byte{})
	err = fdfMerge(bytes.NewReader(pdfData), bytes.NewReader(fdfData), output)
	if err != nil {
		error500Handler(w, r, fmt.Errorf("PDF/FDF Merge Failed: %s", err))
		return
	}
	output.Seek(0, 0)

	w.Header().Add("Content-type", "application/pdf; charset=utf-8")
	w.WriteHeader(200)
	w.Write(output.Bytes())
}

func stripFDF(data []byte) []byte {
	if len(data) > 2 {
		length := len(data)

		if data[0] == []byte("'")[0] {
			data = data[1:]
			length = len(data)
		}

		if data[length-1] == []byte("'")[0] {
			data = data[:length-1]
		}
	}

	return data
}
