package mypdf

type delimiter struct {
	LF  byte
	CR  byte
	EOF [2]byte
}

const (
	pdfNull            byte = 0x00
	pdfLineFeed        byte = 0x0A
	pdfCarriengeReturn byte = 0x0D
	pdfTab             byte = 0x09
	pdfFormFeed        byte = 0x0C
	pdfSpace           byte = 0x20
)

var (
	pdfComment   byte = '%'
	pdfStartXRef      = "startxref"
	pdfXRef           = "xref"
)

type PDFMetadata struct {
	Version        string
	ReferenceTable XReference
}

//////////////////////////////// PDF Objects //////////////////////////////////
type PDFObject struct {
	Offset int
	Raw    []byte
}

//////////////////////////////// xReference //////////////////////////////////
type XReference struct {
	Unknown        int
	NumberOfObject int
	Objects        []*RefObject
}

type RefObject struct {
	Address int
	Unknow  int
	Flag    string
}
