package main

import (
	"fmt"
	"io"

	"github.com/unidoc/unidoc/pdf/annotator"
	"github.com/unidoc/unidoc/pdf/fdf"
	"github.com/unidoc/unidoc/pdf/model"
)

func fdfMerge(pdfInput io.ReadSeeker, fdfInput io.ReadSeeker, pdfOutput io.WriteSeeker) error {
	fdfData, err := fdf.Load(fdfInput)
	if err != nil {
		return fmt.Errorf("Loading FDF: %s", err)
	}

	pdfReader, err := model.NewPdfReader(pdfInput)
	if err != nil {
		return fmt.Errorf("Creating PdfReader: %s", err)
	}

	// Populate the form data.
	err = pdfReader.AcroForm.Fill(fdfData)
	if err != nil {
		return fmt.Errorf("Filling AcroForm: %s", err)
	}

	// Flatten form.
	fieldAppearance := annotator.FieldAppearance{OnlyIfMissing: false}
	err = pdfReader.FlattenFields(true, fieldAppearance)
	if err != nil {
		return fmt.Errorf("Flattening Fields: %s", err)
	}

	// Write out.
	pdfWriter := model.NewPdfWriter()
	pdfWriter.SetVersion(pdfReader.PdfVersionInt())
	pdfWriter.SetForms(nil)

	for i, p := range pdfReader.PageList {
		err := pdfWriter.AddPage(p)
		if err != nil {
			return fmt.Errorf("Adding Page (%d of %d): %s", (i + 1), len(pdfReader.PageList), err)
		}
	}

	err = pdfWriter.Write(pdfOutput)
	if err != nil {
		return fmt.Errorf("Writing PDF: %s", err)
	}
	return nil
}
