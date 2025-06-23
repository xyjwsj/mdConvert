package mdConvert

import (
	parser2 "github.com/xyjwsj/md-parser"
	"os"
	"testing"
)

func TestHtmlRender(t *testing.T) {
	file, _ := os.ReadFile("/Users/wushaojie/Downloads/MDNote.md")
	lexer := parser2.NewLexer(string(file))
	parser := parser2.NewParser(lexer)
	ast := parser.Parse()
	//render := HtmlRender{}
	//NewRender(&render).Render(ast)
	//fmt.Println(render.OutFile(""))

	render := CreatePdfRender()
	NewRender(render).Render(ast)
	render.OutFile("/Users/wushaojie/Downloads/test.pdf")
}
