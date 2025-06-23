package mdConvert

import parser "github.com/xyjwsj/md-parser"

type RenderItem interface {
	RenderTag(node *parser.Node) TagInfo
	RenderText(content string)
	OutFile(path string) string
}
