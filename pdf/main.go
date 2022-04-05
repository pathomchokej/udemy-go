package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number : %g!!!", e)
}

func main() {
	// var w Writer

	// // os.Stdout implements Writer
	// w = os.Stdout
	// //os.Stdout.Read()

	// fmt.Fprintf(w, "hello, writer\n")

	// fmt.Println(ErrNegativeSqrt(-2))

	fmt.Println("Test PDF")

	path, _ := os.Getwd()
	templateFile := path + "/resources/endorsement_original.pdf"
	data, err := os.ReadFile(templateFile)
	if err != nil {
		fmt.Println("Load ", templateFile, " failed :", err)
		return
	}
	fmt.Println()

	targetFile := path + "/temp/endorsement_mod.pdf"
	f, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println("Open ", targetFile, " failed :", err)
		panic(err)
	}
	defer f.Close()

	//n, err := f.Write([]byte("writing some data into a file"))
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println("wrote %d bytes", n)

	fmt.Println("Source : ", templateFile)
	fmt.Println("Target : ", targetFile)

	err = fillOutEndorsementPdf(data, "test", "1.0", f)
	if err != nil {
		fmt.Println("fillOutEndorsementPdf failed :", err)
		panic(err)
	}
}

func convertImageForPdf(signatureS3Bytes []byte, name string, containerX, containerWidth, containerHeight float64) (io.Reader, string, string, float64, error) {

	reader := bytes.NewReader(signatureS3Bytes)
	readerTemp := bytes.NewReader(signatureS3Bytes)
	config, format, err := image.DecodeConfig(readerTemp)
	if err != nil {
		return reader, "", "", 0, err
	}

	imageExt := "png"
	if format == "jpeg" {
		imageExt = "jpg"
	}

	imageName := fmt.Sprintf("%s.%s", name, imageExt)
	imageWidth := containerHeight * float64(config.Width/config.Height)
	imageX := float64(containerX) + float64(containerWidth/2) - float64(imageWidth/2)

	return reader, imageName, imageExt, imageX, nil
}

func fillOutEndorsementPdf(pdfTemplate []byte, effectiveDate, currentVersion string, w io.Writer) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)

	currentPath, _ := os.Getwd()
	currentPath = filepath.Dir(currentPath)
	fontPathR := path.Join(currentPath, "asset/font/THSarabunNew/THSarabunNew.ttf")
	fontPathB := path.Join(currentPath, "asset/font/THSarabunNew/THSarabunNew Bold.ttf")
	fmt.Println("Font R : ", fontPathR)
	fmt.Println("Font B : ", fontPathB)
	// pdf.AddUTF8Font("THSarabunNew", "R", fontPathR)
	// pdf.AddUTF8Font("THSarabunNew", "B", fontPathB)

	pdf.SetTitle("บันทึกสลักหลัง", false)

	// convert []byte to io.ReadSeeker
	rs := io.ReadSeeker(bytes.NewReader(pdfTemplate))

	// create a new Importer instance
	imp := gofpdi.NewImporter()

	// import first page and determine page sizes
	imp.ImportPageFromStream(pdf, &rs, 1, "/MediaBox")
	pageSizes := imp.GetPageSizes()
	pages := len(pageSizes)

	for i := 1; i <= pages; i++ {
		var sizeType gofpdf.SizeType
		sizeType.Wd = pageSizes[i]["/MediaBox"]["w"]
		sizeType.Ht = pageSizes[i]["/MediaBox"]["h"]

		tpl := imp.ImportPageFromStream(pdf, &rs, i, "/MediaBox")

		pdf.AddPageFormat("P", sizeType)
		imp.UseImportedTemplate(pdf, tpl, 0, 0, pageSizes[i]["/MediaBox"]["w"], pageSizes[i]["/MediaBox"]["h"])

		// 	if i == 9 {
		// 		// EffectiveDate
		pdf.SetDrawColor(255, 0, 0)
		pdf.SetLineWidth(1)
		pdf.Rect(217, 620, 255, 12, "D")

		pdf.SetFont("Arial", "B", 40)
		pdf.SetTextColor(255, 0, 0)
		pdf.SetXY(217, 620)
		pdf.CellFormat(255, 12, effectiveDate, "", 0, "C", false, 0, "")

		// 		// Signature1
		// 		pdf.SetDrawColor(255, 0, 0)
		// 		pdf.SetLineWidth(1)
		// 		pdf.Rect(58, 640, 180, 25, "D")

		// 		// sign1S3Key := fmt.Sprintf("%s/%s_sign1.png", s3Config.PdfPath.EndorsementTemplatePath, currentVersion)
		// 		// sign1S3Bytes, err := h.documentServiceRepo.FindS3ObjectByKey(sign1S3Key, s3Config)
		// 		// if err != nil {
		// 		// 	return fmt.Errorf("FindS3ObjectByKey|%s|%s", sign1S3Key, err.Error())
		// 		// }

		// 		// sign1Reader, sign1Name, sign1Ext, sign1X, err := convertImageForPdf(sign1S3Bytes, "sign1", 58, 180, 25)
		// 		// if err != nil {
		// 		// 	return fmt.Errorf("convertImageForPdf|sign1|%s", err.Error())
		// 		// }

		// 		// pdf.RegisterImageOptionsReader(sign1Name, gofpdf.ImageOptions{ImageType: sign1Ext}, sign1Reader)
		// 		// pdf.Image(sign1Name, sign1X, 640, 0, 25, false, "", 0, "")

		// 	}
	}

	// Protect PDF
	//pdf.SetProtection(gofpdf.CnProtectPrint|gofpdf.CnProtectModify, "123456789", "mtlfit")

	if err := pdf.Output(w); err != nil {
		return err
	}
	pdf.Close()

	return nil
}
