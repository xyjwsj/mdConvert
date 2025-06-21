package mdConvert

import (
	"fmt"
	parser2 "github.com/xyjwsj/md-parser"
	"os"
	"testing"
)

func TestHtmlRender(t *testing.T) {
	file, _ := os.ReadFile("/Users/wushaojie/Downloads/MDNote.md")
	lexer := parser2.NewLexer(string(file))
	parser := parser2.NewParser(lexer)
	ast := parser.Parse()
	renderer := NewRenderer()
	html := renderer.Render(ast)

	fmt.Println(html)
}
