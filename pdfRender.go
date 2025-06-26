package mdConvert

import (
	"embed"
	"github.com/phpdave11/gofpdf"
	parser "github.com/xyjwsj/md-parser"
	"log"
	"path/filepath"
	"strings"
)

//go:embed SourceHanSansSC-Normal-Min.ttf SourceHanSansSC-Bold-Min.ttf
var fontAssets embed.FS

//var fontName = "Arial"

var fontName = "SourceHanSansSC"

type PDFRender struct {
	file      *gofpdf.Fpdf
	line      float64
	strong    bool
	cellWidth float64
	imageDir  string
}

var fonts = map[string]string{
	"":  "SourceHanSansSC-Normal-Min.ttf",
	"B": "SourceHanSansSC-Bold-Min.ttf",
}

func CreatePdfRender() *PDFRender {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 15, 15)
	pdf.AddPage()

	for k, v := range fonts {
		file, err := fontAssets.ReadFile(v)
		if err != nil {
			log.Println(err)
		}
		pdf.AddUTF8FontFromBytes(fontName, k, file)
	}

	pdf.SetFont(fontName, "", 14)

	return &PDFRender{
		file:   pdf,
		strong: false,
	}
}

func (pdf *PDFRender) SetImageDir(dir string) {
	pdf.imageDir = dir
}

func (pdf *PDFRender) resetFont() {
	pdf.file.SetFont(fontName, "", 14)
}

func (pdf *PDFRender) RenderTag(node *parser.Node) TagInfo {
	// 自动分页检测
	if pdf.file.GetY() > 280 {
		//pdf.file.AddPage()
	}

	switch node.Type {
	case parser.TokenHeader:
		pdf.file.SetFont(fontName, "B", float64(18-node.Level))
		pdf.line = 10
		if node.Level > 0 {
			pdf.line = 8
		}
		pdf.file.Ln(pdf.line) // 加大标题与正文之间的间距
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
		pdf.line = float64(7 - node.Indent)

		start := "●"
		if node.Indent > 0 {
			start = "○"
		}

		return TagInfo{
			StartFormat: start,
			End:         "\n",
		}

	case parser.TokenEmphasis:
		pdf.file.SetFont(fontName, "", 12)
		pdf.file.SetFillColor(240, 240, 240) // 浅灰色背景
		pdf.line = 5

	case parser.TokenStrong:
		pdf.file.SetFontStyle("B")
		pdf.line = 5
		pdf.strong = true

	case parser.TokenCodeBlock:
		pdf.file.SetFont(fontName, "", 12)
		pdf.file.SetFillColor(240, 240, 240) // 浅灰色背景
		pdf.file.SetX(20)                    // 缩进 20mm
		pdf.file.SetLeftMargin(10)
		pdf.line = 6
		pdf.file.Ln(2)

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
		pdf.file.SetX(30)
		pdf.cellWidth = float64((210.0 - 60) / len(node.Children))
		return TagInfo{
			StartFormat: "",
			End:         "\n",
		}

	case parser.TokenTableCell:
		pdf.line = 6

	}

	return TagInfo{}
}
func (pdf *PDFRender) RenderText(tType parser.TokenType, content, link string) {
	if tType == parser.TokenTableCell {
		if content != "" {
			pdf.file.CellFormat(pdf.cellWidth, pdf.line, content, "1", 0, "L", false, 0, "")
		}
	} else if tType == parser.TokenCodeBlock {
		//log.Println(content)
		pdf.file.MultiCell(0, pdf.line, content, "", "", true)
	} else if tType == parser.TokenEmphasis {
		pdf.file.MultiCell(0, pdf.line, content, "", "", true)
	} else if tType == parser.TokenImage {
		base := filepath.Base(link)
		pdf.file.Image(filepath.Join(pdf.imageDir, base), 0, 0, 0, 0, true, "", 0, "")
	} else {
		//pdf.file.Write(pdf.line, content)
		pdf.multiContentWrite(content)
	}
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

func (pdf *PDFRender) multiContentWrite(content string) {
	split := strings.Split(content, "")
	for _, item := range split {
		if item == "●" {
			pdf.draw(true)
			continue
		}
		if item == "○" {
			pdf.draw(false)
			continue
		}
		pdf.file.Write(pdf.line, item)
	}
}

func (pdf *PDFRender) draw(fill bool) {
	// 获取当前位置
	x := pdf.file.GetX()
	y := pdf.file.GetY()

	radius := 0.5
	pdf.file.SetDrawColor(0, 0, 0) // 黑色边框
	if fill {
		// 实心圆圈（●）
		pdf.file.SetFillColor(0, 0, 0)                     // 黑色填充
		pdf.file.Ellipse(x+2, y+2, radius, radius, 0, "F") // 填充
		pdf.file.SetX(x + 6)                               // 向右偏移一点
	} else {
		// 空心圆圈（○）
		pdf.file.SetFillColor(255, 255, 255)               // 白色填充
		pdf.file.Ellipse(x+2, y+2, radius, radius, 0, "D") // 描边
		pdf.file.SetX(x + 6)
	}
}
