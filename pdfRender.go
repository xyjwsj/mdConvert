package mdConvert

import (
	"embed"
	"github.com/phpdave11/gofpdf"
	parser "github.com/xyjwsj/md-parser"
	"log"
)

//go:embed SFNSRounded.ttf
var fontAssets embed.FS

var fontName = "SFNSRounded"

type PDFRender struct {
	file   *gofpdf.Fpdf
	line   float64
	strong bool
}

func CreatePdfRender() *PDFRender {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 15, 15)
	pdf.AddPage()

	file, _ := fontAssets.ReadFile("SFNSRounded.ttf")
	pdf.AddUTF8FontFromBytes(fontName, "", file)
	pdf.SetFont(fontName, "", 14)

	return &PDFRender{
		file:   pdf,
		strong: false,
	}
}

func (pdf *PDFRender) resetFont() {
	pdf.file.SetFont(fontName, "", 14)
}

func (pdf *PDFRender) RenderTag(node *parser.Node) TagInfo {
	// 自动分页检测
	if pdf.file.GetY() > 280 {
		pdf.file.AddPage()
	}

	switch node.Type {
	case parser.TokenHeader:
		pdf.file.SetFont(fontName, "B", float64(16-node.Level))
		pdf.line = 8
		line := 8
		if node.Level > 0 {
			//line = 4
			pdf.line = 5
		}
		pdf.file.Ln(float64(line)) // 加大标题与正文之间的间距

	case parser.TokenParagraph:
		pdf.file.SetFont(fontName, "", 12)
		pdf.line = 5
		pdf.file.Ln(3)

	case parser.TokenList:
		pdf.file.SetFont(fontName, "", 12)
		pdf.file.Ln(3)
		pdf.line = 6

	case parser.TokenListItem:
		pdf.file.SetFont(fontName, "", 12)
		pdf.file.SetX(20) // 缩进 20mm
		pdf.file.Ln(float64(6 - node.Indent))
		pdf.line = float64(5 - node.Indent)
		start := "• "
		if node.Indent > 0 {
			start = "◦ "
		}
		//pdf.file.Write(4, start)
		return TagInfo{
			StartFormat: start,
			End:         "\n",
		}

	case parser.TokenEmphasis:
		pdf.file.SetFont(fontName, "I", 12)
		pdf.line = 5

	case parser.TokenStrong:
		pdf.file.SetFont(fontName, "B", 12)
		pdf.line = 5
		pdf.strong = true

	case parser.TokenCodeBlock:
		pdf.file.SetFont("Courier", "", 12)
		pdf.file.SetFillColor(240, 240, 240) // 浅灰色背景
		pdf.line = 5

	case parser.TokenHorizontalRule:
		x, y := pdf.file.GetX(), pdf.file.GetY()
		pdf.file.Line(x, y, x+180, y)
		pdf.file.Ln(3)

	case parser.TokenLink:
		pdf.file.SetFont(fontName, "U", 12)
		pdf.line = 5

	case parser.TokenImage:
		// 插入图片，宽度固定为 50mm，高度自动计算
		pdf.file.ImageOptions(node.Content, pdf.file.GetX(), pdf.file.GetY(), 50, 0, false, gofpdf.ImageOptions{}, 0, "")
		pdf.file.Ln(10)
		pdf.line = 10

	case parser.TokenTable:
		pdf.file.SetFont(fontName, "", 12)
		pdf.line = 6
		pdf.file.Ln(3)

	case parser.TokenTableRow:
		pdf.file.SetFont(fontName, "", 12)
		//pdf.file.Ln(1)
		pdf.line = 6

	case parser.TokenTableCell:
		pdf.line = 6

	}

	return TagInfo{}
}
func (pdf *PDFRender) RenderText(tType parser.TokenType, content string) {
	// pdf.file.MultiCell(0, pdf.line, content, "", "", false)
	if tType == parser.TokenTableRow {
		pdf.file.CellFormat(0, pdf.line, content, "", 0, "", false, 0, "")
		//pdf.file.CellFormat(0, pdf.line, start+" ", "", 0, "", false, 0, "")
	}
	pdf.file.Write(pdf.line, content)
	if content != "" && pdf.strong && tType == parser.TokenText {
		pdf.resetFont()
	}
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

func (pdf *PDFRender) styleConfig(t parser.TokenType) {

}
