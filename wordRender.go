package mdConvert

import (
	"github.com/xyjwsj/godocx"
	"github.com/xyjwsj/godocx/docx"
	parser "github.com/xyjwsj/md-parser"
	"log"
	"path/filepath"
	"strings"
)

type WordRender struct {
	doc        *docx.RootDoc
	paragraph  *docx.Paragraph
	parentType parser.TokenType
	textType   parser.TokenType
	level      int
	indent     int
	table      *docx.Table
	row        *docx.Row

	imageDir    string
	contextPath string
}

func CreateWordRender() *WordRender {
	// Create New Document
	document, err := godocx.NewDocument()
	if err != nil {
		log.Fatal(err)
	}
	return &WordRender{
		doc: document,
	}
}

func (word *WordRender) SetImageDir(dir string) {
	word.imageDir = dir
}

func (word *WordRender) SetContentPath(path string) {
	word.contextPath = path
}

func createTab(num int) string {
	builder := strings.Builder{}
	for i := 0; i < num; i++ {
		builder.WriteString("    ")
	}
	return builder.String()
}

func (word *WordRender) RenderTag(node *parser.Node) TagInfo {
	if node.Type != parser.TokenText {
		word.textType = parser.TokenEOF
	}
	switch node.Type {
	case parser.TokenHeader:
		word.parentType = node.Type
		word.level = node.Level
		word.paragraph = word.doc.AddParagraph("")
	case parser.TokenParagraph:
		word.parentType = node.Type
		word.paragraph = word.doc.AddParagraph("")
	case parser.TokenList:
	case parser.TokenListItem:
		flag := "•"
		if node.Indent > 0 {
			flag = "◦"
		}
		word.parentType = node.Type
		word.paragraph = word.doc.AddParagraph(createTab(node.Indent) + flag + "  ")
	case parser.TokenEmphasis:

	case parser.TokenStrong:
		word.textType = parser.TokenStrong
	case parser.TokenCodeBlock:
		//word.paragraph = word.doc.AddParagraph("")

	case parser.TokenHorizontalRule:
		word.paragraph = word.doc.AddParagraph("")

	case parser.TokenLink:
		word.paragraph = word.doc.AddParagraph("")

	case parser.TokenImage:
		//word.paragraph = word.doc.AddParagraph("")

	case parser.TokenTable:
		word.parentType = node.Type
		word.table = word.doc.AddTable()
		word.table.Style("LightList-Accent4")

	case parser.TokenTableRow:
		word.parentType = node.Type
		word.row = word.table.AddRow()

	case parser.TokenTableCell:
		word.parentType = node.Type
	}

	return TagInfo{}
}
func (word *WordRender) RenderText(tType parser.TokenType, content, link string) {

	if tType == parser.TokenText {
		if word.paragraph != nil {
			switch word.parentType {
			case parser.TokenHeader:
				word.doc.AddHeading(content, uint(word.level))
			case parser.TokenList:
			case parser.TokenEmphasis:

			case parser.TokenCodeBlock:

			case parser.TokenHorizontalRule:

			case parser.TokenLink:

			case parser.TokenImage:
			default:
				text := word.paragraph.AddText(content)
				text.Bold(word.textType == parser.TokenStrong)
			}
		}

		word.parentType = parser.TokenEOF
		word.textType = parser.TokenEOF
	} else {
		switch tType {
		case parser.TokenLink:

		case parser.TokenImage:
			base := filepath.Base(link)
			if word.contextPath != "" {
				base = strings.ReplaceAll(link, word.contextPath, "")
			}
			imgPath := filepath.Join(word.imageDir, base)
			fType, _ := DetectImageFormat(imgPath)
			_, err := word.paragraph.AddPicture(imgPath, strings.ToLower(fType), 100, 100)
			log.Println(err)
		case parser.TokenTable:
		case parser.TokenTableCell:
			cell := word.row.AddCell()
			cell.AddParagraph(content)
		case parser.TokenCodeBlock:
			split := strings.Split(content, "\n")
			for _, itm := range split {
				word.doc.AddParagraph(itm)
			}
		default:
			text := word.paragraph.AddText(content)
			text.Bold(word.textType == parser.TokenStrong)
		}
	}
}
func (word *WordRender) OutFile(path string) string {
	err := word.doc.SaveTo(path)
	if err != nil {
		log.Println(err)
		return ""
	}
	return ""
}
