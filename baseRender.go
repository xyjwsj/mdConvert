package mdConvert

import (
	parser "github.com/xyjwsj/md-parser"
	"html"
)

// Render 渲染器
type Render struct {
	item RenderItem
}

// NewRender 创建一个新的渲染器
func NewRender(item RenderItem) *Render {
	return &Render{
		item: item,
	}
}

// Render 渲染AST为自定义渲染器
func (r *Render) Render(node *parser.Node) string {
	r.renderLine(node, 0)
	return r.item.Result()
}

func (r *Render) renderLine(node *parser.Node, line int) {

	tagInfo := r.item.RenderTag(node)
	r.item.RenderText(createSpace(node.Indent) + tagInfo.StartFormat)
	if node.Children != nil && len(node.Children) > 0 {
		for _, child := range node.Children {
			r.renderLine(child, line+1)
		}
	} else {
		r.item.RenderText(createSpace(node.Indent) + html.EscapeString(node.Content))
	}

	r.item.RenderText(tagInfo.End)
	if line < 2 {
		r.item.RenderText("\n")
	}
}

func createSpace(len int) string {
	str := ""
	for i := 0; i < len; i++ {
		str += " "
	}
	return str
}
