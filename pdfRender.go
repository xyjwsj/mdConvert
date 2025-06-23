package mdConvert

import (
	"github.com/phpdave11/gofpdf"
	parser "github.com/xyjwsj/md-parser"
	"log"
)

type PDFRender struct {
	file *gofpdf.Fpdf
	line float64
}

func CreatePdfRender() *PDFRender {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 15, 15)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	return &PDFRender{
		file: pdf,
	}
}

func (pdf *PDFRender) RenderTag(node *parser.Node) TagInfo {
	switch node.Type {
	case parser.TokenHeader:
		pdf.file.SetFont("Arial", "B", float64(18-node.Level))
		pdf.line = 6
		pdf.file.Ln(3)
	case parser.TokenParagraph:
		pdf.file.SetFont("Arial", "", 24)
		pdf.file.Ln(3)
	case parser.TokenList:
		pdf.file.SetFont("Arial", "", 24)
		pdf.file.Ln(3)
	case parser.TokenListItem:
		pdf.file.SetFont("Arial", "", 24)
	case parser.TokenEmphasis:
		pdf.file.SetFont("Arial", "", 24)
	case parser.TokenStrong:
		pdf.file.SetFont("Arial", "", 24)
	case parser.TokenCodeBlock:
		pdf.file.SetFont("Arial", "", 24)
	case parser.TokenHorizontalRule:
		pdf.file.SetFont("Arial", "", 24)
	case parser.TokenLink:
		pdf.file.SetFont("Arial", "", 24)
	case parser.TokenImage:
		pdf.file.SetFont("Arial", "", 24)
	case parser.TokenTable:
		pdf.file.SetFont("Arial", "", 24)
	case parser.TokenTableRow:
		pdf.file.SetFont("Arial", "", 24)
	case parser.TokenTableCell:
		pdf.file.SetFont("Arial", "", 24)
	}
	return TagInfo{}
}
func (pdf *PDFRender) RenderText(content string) {
	pdf.file.MultiCell(0, pdf.line, content, "", "", false)
}
func (pdf *PDFRender) OutFile(path string) string {
	// 输出 PDF
	err := pdf.file.OutputFileAndClose(path)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return ""
}
