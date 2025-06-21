package mdConvert

import parser "github.com/xyjwsj/md-parser"

type Render interface {
	Render(node *parser.Node)
}
