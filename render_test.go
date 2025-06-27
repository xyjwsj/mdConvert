package mdConvert

import (
	parser2 "github.com/xyjwsj/md-parser"
	"os"
	"testing"
)

func TestPdfRender(t *testing.T) {
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

func TestWordRender(t *testing.T) {
	file, _ := os.ReadFile("/Users/wushaojie/Downloads/parse.md")
	lexer := parser2.NewLexer(string(file))
	parser := parser2.NewParser(lexer)
	ast := parser.Parse()
	//render := HtmlRender{}
	//NewRender(&render).Render(ast)
	//fmt.Println(render.OutFile(""))

	render := CreateWordRender()
	render.SetImageDir("/Users/wushaojie/Library/Containers/com.allen.mdnote/Data/Library/Application Support/LiveMark/image")
	//render.SetImageDir("/Users/wushaojie/Downloads")
	render.SetContentPath("/api/resource/")

	newRender := NewRender(render)
	newRender.Render(ast)
	render.OutFile("/Users/wushaojie/Downloads/test.docx")
}
