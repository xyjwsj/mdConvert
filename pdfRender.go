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

	// 添加并使用中文字体
	//pdf.AddFont("Arial", "", "font/Arial.ttf")
	pdf.SetFont("Arial", "", 14)

	return &PDFRender{
		file: pdf,
	}
}

func (pdf *PDFRender) RenderTag(node *parser.Node) TagInfo {
	// 自动分页检测
	if pdf.file.GetY() > 280 {
		pdf.file.AddPage()
	}

	switch node.Type {
	case parser.TokenHeader:
		pdf.file.SetFont("Arial", "B", float64(18-node.Level))
		pdf.file.Ln(5) // 加大标题与正文之间的间距

	case parser.TokenParagraph:
		pdf.file.SetFont("Arial", "", 14)
		pdf.line = 5
		pdf.file.Ln(3)

	case parser.TokenList:
		pdf.file.SetFont("Arial", "", 14)
		pdf.file.Ln(3)
		pdf.line = 5

	case parser.TokenListItem:
		pdf.file.SetFont("Arial", "", 14)
		pdf.file.SetX(20) // 缩进 20mm
		pdf.line = 5

	case parser.TokenEmphasis:
		pdf.file.SetFont("Arial", "I", 14)
		pdf.line = 5

	case parser.TokenStrong:
		pdf.file.SetFont("Arial", "B", 14)
		pdf.line = 5

	case parser.TokenCodeBlock:
		pdf.file.SetFont("Courier", "", 12)
		pdf.file.SetFillColor(240, 240, 240) // 浅灰色背景
		pdf.line = 5

	case parser.TokenHorizontalRule:
		x, y := pdf.file.GetX(), pdf.file.GetY()
		pdf.file.Line(x, y, x+180, y)
		pdf.file.Ln(3)

	case parser.TokenLink:
		pdf.file.SetFont("Arial", "U", 14)
		pdf.line = 5

	case parser.TokenImage:
		// 插入图片，宽度固定为 50mm，高度自动计算
		pdf.file.ImageOptions(node.Content, pdf.file.GetX(), pdf.file.GetY(), 50, 0, false, gofpdf.ImageOptions{}, 0, "")
		pdf.file.Ln(10)
		pdf.line = 10

	case parser.TokenTable:
		pdf.file.SetFont("Arial", "", 14)
		pdf.line = 6

	case parser.TokenTableRow:
		pdf.file.SetFont("Arial", "", 14)
		pdf.file.Ln(1)
		pdf.line = 6

	case parser.TokenTableCell:
		pdf.file.CellFormat(40, 6, node.Content, "1", 0, "L", false, 0, "")
		pdf.line = 6

	case parser.TokenText:
		pdf.file.SetFont("Arial", "", 14)
		pdf.line = 5
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
