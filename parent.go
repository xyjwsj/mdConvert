package mdConvert

import parser "github.com/xyjwsj/md-parser"

type RenderItem interface {
	RenderTag(node *parser.Node) TagInfo
	RenderText(tType parser.TokenType, content string)
	OutFile(path string) string
}
