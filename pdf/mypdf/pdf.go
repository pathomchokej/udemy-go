package mypdf

import (
	"fmt"
	"io/ioutil"
)

type PDF interface {
	GetMetadata() (*PDFMetadata, error)
}

type pdfInformation struct {
	metadata *PDFMetadata
}

type pdfRaw struct {
	startXref int
}

func CreatePDFInstance(filename string) (PDF, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Read %s got error : %s", filename, err)
	}

	// TODO:
	// get pdf metadata
	metadata, err := validatePDFStructure(data)
	if err != nil {
		return nil, fmt.Errorf("Read %s got error : %s", filename, err)
	}

	// load pdf by vesion

	pdfInfo := pdfInformation{
		metadata: metadata,
	}

	return pdfInfo, nil
}

//////////////////////////////// Private //////////////////////////////////////////////////////////////////
func validatePDFStructure(data []byte) (*PDFMetadata, error) {

	return nil, nil
}

//////////////////////////////// Public //////////////////////////////////////////////////////////////////
func (pdf pdfInformation) GetMetadata() (*PDFMetadata, error) {

	return nil, nil
}
