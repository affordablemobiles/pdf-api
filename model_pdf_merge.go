package main

import (
	"io"

	"github.com/unidoc/unidoc/pdf/annotator"
	"github.com/unidoc/unidoc/pdf/fdf"
	"github.com/unidoc/unidoc/pdf/model"
)

func fdfMerge(pdfInput io.ReadSeeker, fdfInput io.ReadSeeker, pdfOutput io.WriteSeeker) error {
	fdfData, err := fdf.Load(fdfInput)
	if err != nil {
		return err
	}

	pdfReader, err := model.NewPdfReader(pdfInput)
	if err != nil {
		return err
	}

	// Populate the form data.
	err = pdfReader.AcroForm.Fill(fdfData)
	if err != nil {
		return err
	}

	// Flatten form.
	fieldAppearance := annotator.FieldAppearance{OnlyIfMissing: false}
	err = pdfReader.FlattenFields(true, fieldAppearance)
	if err != nil {
		return err
	}

	// Write out.
	pdfWriter := model.NewPdfWriter()
	pdfWriter.SetForms(nil)

	for _, p := range pdfReader.PageList {
		err := pdfWriter.AddPage(p)
		if err != nil {
			return err
		}
	}

	err = pdfWriter.Write(pdfOutput)
	return err
}
