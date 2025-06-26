package mdConvert

import (
	parser2 "github.com/xyjwsj/md-parser"
	"os"
	"testing"
)

func TestHtmlRender(t *testing.T) {
	file, _ := os.ReadFile("/Users/wushaojie/Downloads/parse.md")
	lexer := parser2.NewLexer(string(file))
	parser := parser2.NewParser(lexer)
	ast := parser.Parse()
	//render := HtmlRender{}
	//NewRender(&render).Render(ast)
	//fmt.Println(render.OutFile(""))

	render := CreatePdfRender()
	render.SetImageDir("/Users/wushaojie/Library/Containers/com.allen.mdnote/Data/Library/Application Support/LiveMark/image")
	render.SetContentPath("/api/resource/")
	newRender := NewRender(render)
	newRender.Render(ast)
	render.OutFile("/Users/wushaojie/Downloads/test.pdf")
}
